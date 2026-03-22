package channels

import (
	"encoding/base64"
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/4claw/4claw/pkg/bus"
	"github.com/4claw/4claw/pkg/config"
	"github.com/4claw/4claw/pkg/logger"
)

type WeChatChannel struct {
	*BaseChannel

	cfg       config.WeChatConfig
	account   *resolvedWeChatAccount
	api       *weChatAPIClient
	ctx       context.Context
	cancel    context.CancelFunc
	wg        sync.WaitGroup
	tokenMu   sync.RWMutex
	ctxTokens map[string]string
}

func NewWeChatChannel(cfg config.WeChatConfig, messageBus *bus.MessageBus) (*WeChatChannel, error) {
	resolved, err := resolveWeChatAccount(cfg)
	if err != nil {
		return nil, err
	}
	ch := &WeChatChannel{
		BaseChannel: NewBaseChannel("wechat", cfg, messageBus, cfg.AllowFrom),
		cfg:         cfg,
		account:     resolved,
		api:         newWeChatAPIClient(resolved.BaseURL, resolved.Token, 15*time.Second),
		ctxTokens:   make(map[string]string),
	}
	return ch, nil
}

func (c *WeChatChannel) Start(ctx context.Context) error {
	if c.IsRunning() {
		return nil
	}
	if c.account == nil || c.account.Token == "" {
		return fmt.Errorf("wechat channel is not configured: missing account token")
	}
	c.ctx, c.cancel = context.WithCancel(ctx)
	c.setRunning(true)
	c.wg.Add(1)
	go c.pollLoop()
	logger.InfoCF("wechat", "WeChat channel started", map[string]any{
		"account_id": c.account.AccountID,
		"base_url":   c.account.BaseURL,
	})
	return nil
}

func (c *WeChatChannel) Stop(ctx context.Context) error {
	if !c.IsRunning() {
		return nil
	}
	if c.cancel != nil {
		c.cancel()
	}
	done := make(chan struct{})
	go func() {
		defer close(done)
		c.wg.Wait()
	}()
	select {
	case <-done:
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(5 * time.Second):
		return fmt.Errorf("timeout waiting for wechat channel to stop")
	}
	c.setRunning(false)
	logger.InfoC("wechat", "WeChat channel stopped")
	return nil
}

func (c *WeChatChannel) Send(ctx context.Context, msg bus.OutboundMessage) error {
	if !c.IsRunning() {
		return fmt.Errorf("wechat channel not running")
	}
	contextToken := c.getContextToken(msg.ChatID)
	if contextToken == "" {
		return fmt.Errorf("wechat send requires context token for chat %s", msg.ChatID)
	}

	if msg.Metadata != nil && msg.Metadata["account_id"] != "" && msg.Metadata["account_id"] != c.account.AccountID {
		return fmt.Errorf("wechat outbound account mismatch: %s != %s", msg.Metadata["account_id"], c.account.AccountID)
	}

	if strings.TrimSpace(msg.Media) == "" {
		return c.sendText(ctx, msg.ChatID, contextToken, msg.Content)
	}
	return c.sendMedia(ctx, msg.ChatID, contextToken, msg.Content, msg.Media)
}

func (c *WeChatChannel) pollLoop() {
	defer c.wg.Done()
	cursor := loadWeChatSyncBuf(c.account.AccountID)
	timeoutMS := c.cfg.LongPollTimeoutMS
	if timeoutMS <= 0 {
		timeoutMS = 35000
	}
	failures := 0
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
		}

		resp, err := c.api.getUpdates(c.ctx, cursor, timeoutMS)
		if err != nil {
			failures++
			logger.ErrorCF("wechat", "getUpdates failed", map[string]any{
				"account_id": c.account.AccountID,
				"error":      err.Error(),
				"failures":   failures,
			})
			select {
			case <-time.After(2 * time.Second):
				continue
			case <-c.ctx.Done():
				return
			}
		}
		failures = 0
		if resp.LongPollingTimeoutMS > 0 {
			timeoutMS = resp.LongPollingTimeoutMS
		}
		if resp.GetUpdatesBuf != "" && resp.GetUpdatesBuf != cursor {
			cursor = resp.GetUpdatesBuf
			if err := saveWeChatSyncBuf(c.account.AccountID, cursor); err != nil {
				logger.WarnCF("wechat", "Failed to persist get_updates_buf", map[string]any{
					"account_id": c.account.AccountID,
					"error":      err.Error(),
				})
			}
		}
		if resp.ErrCode == -14 {
			logger.ErrorCF("wechat", "WeChat session expired", map[string]any{
				"account_id": c.account.AccountID,
			})
			select {
			case <-time.After(30 * time.Second):
				continue
			case <-c.ctx.Done():
				return
			}
		}
		for _, inbound := range resp.Messages {
			c.handleInbound(c.ctx, inbound)
		}
	}
}

func (c *WeChatChannel) handleInbound(ctx context.Context, inbound weChatMessage) {
	fromUserID := strings.TrimSpace(inbound.FromUserID)
	if fromUserID == "" {
		return
	}
	if inbound.ContextToken != "" {
		c.setContextToken(fromUserID, inbound.ContextToken)
	}
	content := extractWeChatMessageText(inbound.ItemList)
	mediaPaths := make([]string, 0, 1)
	if mediaPath, err := c.downloadInboundMedia(ctx, inbound); err != nil {
		logger.WarnCF("wechat", "Failed to download inbound media", map[string]any{
			"account_id": c.account.AccountID,
			"from_user":  fromUserID,
			"error":      err.Error(),
		})
	} else if mediaPath != "" {
		mediaPaths = append(mediaPaths, mediaPath)
	}

	metadata := map[string]string{
		"account_id": c.account.AccountID,
		"peer_kind":  "direct",
		"peer_id":    fromUserID,
	}
	c.HandleMessage(fromUserID, fromUserID, content, mediaPaths, metadata)
}

func (c *WeChatChannel) sendText(ctx context.Context, toUserID, contextToken, text string) error {
	clientID := fmt.Sprintf("wechat-%d", time.Now().UnixNano())
	return c.api.sendMessage(ctx, weChatMessage{
		ToUserID:     toUserID,
		ClientID:     clientID,
		MessageType:  weChatMessageTypeBot,
		MessageState: weChatMessageStateDone,
		ItemList: []weChatMessageItem{
			{
				Type:     weChatMessageTypeText,
				TextItem: &weChatTextItem{Text: text},
			},
		},
		ContextToken: contextToken,
	})
}

func (c *WeChatChannel) sendMedia(ctx context.Context, toUserID, contextToken, text, mediaRef string) error {
	filePath, err := maybeDownloadRemoteMedia(ctx, mediaRef)
	if err != nil {
		return err
	}
	ext := strings.ToLower(filepath.Ext(filePath))
	mediaType := weChatUploadMediaFile
	switch {
	case isImageExt(ext):
		mediaType = weChatUploadMediaImage
	case isVideoExt(ext):
		mediaType = weChatUploadMediaVideo
	}
	uploaded, err := uploadWeChatMediaFile(ctx, c.api, c.account.CDNBaseURL, filePath, toUserID, mediaType)
	if err != nil {
		return err
	}
	if strings.TrimSpace(text) != "" {
		if err := c.sendText(ctx, toUserID, contextToken, text); err != nil {
			return err
		}
	}
	clientID := fmt.Sprintf("wechat-%d", time.Now().UnixNano())
	return c.api.sendMessage(ctx, weChatMessage{
		ToUserID:     toUserID,
		ClientID:     clientID,
		MessageType:  weChatMessageTypeBot,
		MessageState: weChatMessageStateDone,
		ItemList: []weChatMessageItem{
			mediaItemFromUpload(filePath, uploaded),
		},
		ContextToken: contextToken,
	})
}

func (c *WeChatChannel) downloadInboundMedia(ctx context.Context, inbound weChatMessage) (string, error) {
	for _, item := range inbound.ItemList {
		switch item.Type {
		case weChatMessageTypeImage:
			if item.ImageItem == nil || item.ImageItem.Media == nil || item.ImageItem.Media.EncryptQueryParam == "" {
				continue
			}
			aesKey := item.ImageItem.Media.AESKey
			if item.ImageItem.RawAESKey != "" {
				aesKey = encodeHexKeyAsBase64(item.ImageItem.RawAESKey)
			}
			if aesKey == "" {
				continue
			}
			raw, err := downloadAndDecryptWeChatMedia(ctx, c.account.CDNBaseURL, item.ImageItem.Media.EncryptQueryParam, aesKey)
			if err != nil {
				return "", err
			}
			return saveInboundWeChatMedia(raw, ".img")
		case weChatMessageTypeVideo:
			if item.VideoItem == nil || item.VideoItem.Media == nil || item.VideoItem.Media.EncryptQueryParam == "" || item.VideoItem.Media.AESKey == "" {
				continue
			}
			raw, err := downloadAndDecryptWeChatMedia(ctx, c.account.CDNBaseURL, item.VideoItem.Media.EncryptQueryParam, item.VideoItem.Media.AESKey)
			if err != nil {
				return "", err
			}
			return saveInboundWeChatMedia(raw, ".mp4")
		case weChatMessageTypeFile:
			if item.FileItem == nil || item.FileItem.Media == nil || item.FileItem.Media.EncryptQueryParam == "" || item.FileItem.Media.AESKey == "" {
				continue
			}
			raw, err := downloadAndDecryptWeChatMedia(ctx, c.account.CDNBaseURL, item.FileItem.Media.EncryptQueryParam, item.FileItem.Media.AESKey)
			if err != nil {
				return "", err
			}
			suffix := filepath.Ext(item.FileItem.FileName)
			if suffix == "" {
				suffix = ".bin"
			}
			return saveInboundWeChatMedia(raw, suffix)
		case weChatMessageTypeVoice:
			if item.VoiceItem == nil || item.VoiceItem.Media == nil || item.VoiceItem.Media.EncryptQueryParam == "" || item.VoiceItem.Media.AESKey == "" {
				if item.VoiceItem != nil && strings.TrimSpace(item.VoiceItem.Text) != "" {
					return "", nil
				}
				continue
			}
			raw, err := downloadAndDecryptWeChatMedia(ctx, c.account.CDNBaseURL, item.VoiceItem.Media.EncryptQueryParam, item.VoiceItem.Media.AESKey)
			if err != nil {
				return "", err
			}
			return saveInboundWeChatMedia(raw, ".silk")
		}
	}
	return "", nil
}

func extractWeChatMessageText(items []weChatMessageItem) string {
	for _, item := range items {
		switch item.Type {
		case weChatMessageTypeText:
			if item.TextItem != nil && strings.TrimSpace(item.TextItem.Text) != "" {
				if item.RefMsg != nil {
					refText := strings.TrimSpace(extractRefText(item.RefMsg))
					if refText != "" {
						return "[引用: " + refText + "]\n" + item.TextItem.Text
					}
				}
				return item.TextItem.Text
			}
		case weChatMessageTypeVoice:
			if item.VoiceItem != nil && strings.TrimSpace(item.VoiceItem.Text) != "" {
				return item.VoiceItem.Text
			}
		}
	}
	return ""
}

func extractRefText(ref *weChatRefMessage) string {
	if ref == nil {
		return ""
	}
	if ref.Title != "" {
		return ref.Title
	}
	if ref.MessageItem != nil && ref.MessageItem.TextItem != nil {
		return ref.MessageItem.TextItem.Text
	}
	return ""
}

func (c *WeChatChannel) setContextToken(userID, contextToken string) {
	c.tokenMu.Lock()
	defer c.tokenMu.Unlock()
	c.ctxTokens[userID] = contextToken
}

func (c *WeChatChannel) getContextToken(userID string) string {
	c.tokenMu.RLock()
	defer c.tokenMu.RUnlock()
	return c.ctxTokens[userID]
}

func encodeHexKeyAsBase64(hexKey string) string {
	raw, err := decodeHexString(hexKey)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(raw)
}
