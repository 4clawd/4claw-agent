package channels

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/open-dingtalk/dingtalk-stream-sdk-go/chatbot"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/client"

	"github.com/4claw/4claw/pkg/bus"
	"github.com/4claw/4claw/pkg/config"
	"github.com/4claw/4claw/pkg/logger"
	"github.com/4claw/4claw/pkg/utils"
)

// DingTalkChannel implements the Channel interface for DingTalk.
// It uses WebSocket for receiving messages via stream mode and API for sending.
type DingTalkChannel struct {
	*BaseChannel
	config       config.DingTalkConfig
	clientID     string
	clientSecret string
	streamClient *client.StreamClient
	ctx          context.Context
	cancel       context.CancelFunc
	// Map to store session webhooks for each chat.
	sessionWebhooks sync.Map // chatID -> sessionWebhook
	tokenMu         sync.Mutex
	tokenCache      map[string]dingtalkAccessToken
}

type dingtalkAccessToken struct {
	Token     string
	ExpiresAt time.Time
}

// NewDingTalkChannel creates a new DingTalk channel instance.
func NewDingTalkChannel(cfg config.DingTalkConfig, messageBus *bus.MessageBus) (*DingTalkChannel, error) {
	if cfg.ClientID == "" || cfg.ClientSecret == "" {
		return nil, fmt.Errorf("dingtalk client_id and client_secret are required")
	}

	base := NewBaseChannel("dingtalk", cfg, messageBus, cfg.AllowFrom)

	return &DingTalkChannel{
		BaseChannel:  base,
		config:       cfg,
		clientID:     cfg.ClientID,
		clientSecret: cfg.ClientSecret,
		tokenCache:   make(map[string]dingtalkAccessToken),
	}, nil
}

// Start initializes the DingTalk channel with Stream Mode.
func (c *DingTalkChannel) Start(ctx context.Context) error {
	logger.InfoC("dingtalk", "Starting DingTalk channel (Stream Mode)...")

	c.ctx, c.cancel = context.WithCancel(ctx)

	cred := client.NewAppCredentialConfig(c.clientID, c.clientSecret)
	c.streamClient = client.NewStreamClient(
		client.WithAppCredential(cred),
		client.WithAutoReconnect(true),
	)
	c.streamClient.RegisterChatBotCallbackRouter(c.onChatBotMessageReceived)

	if err := c.streamClient.Start(c.ctx); err != nil {
		return fmt.Errorf("failed to start stream client: %w", err)
	}

	c.setRunning(true)
	logger.InfoC("dingtalk", "DingTalk channel started (Stream Mode)")
	return nil
}

// Stop gracefully stops the DingTalk channel.
func (c *DingTalkChannel) Stop(ctx context.Context) error {
	logger.InfoC("dingtalk", "Stopping DingTalk channel...")

	if c.cancel != nil {
		c.cancel()
	}

	if c.streamClient != nil {
		c.streamClient.Close()
	}

	c.setRunning(false)
	logger.InfoC("dingtalk", "DingTalk channel stopped")
	return nil
}

// Send sends a message to DingTalk via the chatbot reply API.
func (c *DingTalkChannel) Send(ctx context.Context, msg bus.OutboundMessage) error {
	if !c.IsRunning() {
		return fmt.Errorf("dingtalk channel not running")
	}

	sessionWebhookRaw, ok := c.sessionWebhooks.Load(msg.ChatID)
	if !ok {
		return fmt.Errorf("no session_webhook found for chat %s, cannot send message", msg.ChatID)
	}

	sessionWebhook, ok := sessionWebhookRaw.(string)
	if !ok {
		return fmt.Errorf("invalid session_webhook type for chat %s", msg.ChatID)
	}

	logger.DebugCF("dingtalk", "Sending message", map[string]any{
		"chat_id": msg.ChatID,
		"preview": utils.Truncate(msg.Content, 100),
	})

	return c.SendDirectReply(ctx, sessionWebhook, msg.Content)
}

// onChatBotMessageReceived is called by the Stream SDK when a new message arrives.
func (c *DingTalkChannel) onChatBotMessageReceived(
	ctx context.Context,
	data *chatbot.BotCallbackDataModel,
) ([]byte, error) {
	if data == nil {
		return nil, nil
	}

	content := strings.TrimSpace(data.Text.Content)
	eventMap := structToMap(data)
	contentMap := normalizeMap(data.Content)
	msgType := firstNonEmpty(
		stringValue(eventMap["msgtype"]),
		stringValue(eventMap["msgType"]),
	)
	if msgType == "" && content != "" {
		msgType = "text"
	}

	corpID := firstNonEmpty(
		stringValue(eventMap["chatbotCorpId"]),
		stringValue(eventMap["senderCorpId"]),
		stringValue(eventMap["corpId"]),
	)

	contentParts := make([]string, 0, 2)
	if content != "" {
		contentParts = append(contentParts, content)
	}
	attachmentText, mediaPaths := c.extractAttachments(ctx, msgType, corpID, contentMap)
	if attachmentText != "" {
		contentParts = append(contentParts, attachmentText)
	}

	content = strings.TrimSpace(strings.Join(contentParts, "\n"))
	if content == "" && len(mediaPaths) == 0 {
		return nil, nil
	}
	if content == "" {
		content = "[attachment message]"
	}

	senderID := data.SenderStaffId
	senderNick := data.SenderNick
	chatID := senderID
	if data.ConversationType != "1" {
		chatID = data.ConversationId
	}

	c.sessionWebhooks.Store(chatID, data.SessionWebhook)

	metadata := map[string]string{
		"sender_name":       senderNick,
		"conversation_id":   data.ConversationId,
		"conversation_type": data.ConversationType,
		"platform":          "dingtalk",
		"session_webhook":   data.SessionWebhook,
		"msg_type":          msgType,
	}
	if corpID != "" {
		metadata["corp_id"] = corpID
	}

	if data.ConversationType == "1" {
		metadata["peer_kind"] = "direct"
		metadata["peer_id"] = senderID
	} else {
		metadata["peer_kind"] = "group"
		metadata["peer_id"] = data.ConversationId
	}

	logger.DebugCF("dingtalk", "Received message", map[string]any{
		"sender_nick": senderNick,
		"sender_id":   senderID,
		"msg_type":    msgType,
		"media_count": len(mediaPaths),
		"preview":     utils.Truncate(content, 50),
	})

	c.HandleMessage(senderID, chatID, content, mediaPaths, metadata)
	return nil, nil
}

// SendDirectReply sends a direct reply using the session webhook.
func (c *DingTalkChannel) SendDirectReply(ctx context.Context, sessionWebhook, content string) error {
	replier := chatbot.NewChatbotReplier()
	contentBytes := []byte(content)
	titleBytes := []byte("4claw")

	err := replier.SimpleReplyMarkdown(
		ctx,
		sessionWebhook,
		titleBytes,
		contentBytes,
	)
	if err != nil {
		return fmt.Errorf("failed to send reply: %w", err)
	}

	return nil
}

func (c *DingTalkChannel) extractAttachments(
	ctx context.Context,
	msgType, corpID string,
	contentMap map[string]any,
) (string, []string) {
	switch strings.ToLower(msgType) {
	case "picture", "image":
		return c.handleDingTalkDownloadCode(ctx, corpID, contentMap, []string{
			"pictureDownloadCode", "downloadCode",
		}, "dingtalk-image", ".png", "[image]")
	case "file":
		fileName := firstNonEmpty(
			stringValue(contentMap["fileName"]),
			stringValue(contentMap["name"]),
			"dingtalk-file",
		)
		path := c.downloadMessageFile(ctx, corpID, firstNonEmpty(
			stringValue(contentMap["downloadCode"]),
			stringValue(contentMap["fileDownloadCode"]),
		), fileName)
		if path == "" {
			return "[file]", nil
		}
		return fmt.Sprintf("[file: %s]", filepath.Base(path)), []string{path}
	case "audio":
		return c.handleDingTalkDownloadCode(ctx, corpID, contentMap, []string{
			"downloadCode", "audioDownloadCode",
		}, "dingtalk-audio", ".mp3", "[audio]")
	case "video":
		return c.handleDingTalkDownloadCode(ctx, corpID, contentMap, []string{
			"downloadCode", "videoDownloadCode",
		}, "dingtalk-video", ".mp4", "[video]")
	case "richtext", "rich_text":
		return c.extractRichTextAttachments(ctx, corpID, contentMap)
	default:
		return "", nil
	}
}

func (c *DingTalkChannel) extractRichTextAttachments(
	ctx context.Context,
	corpID string,
	contentMap map[string]any,
) (string, []string) {
	items, ok := contentMap["richText"].([]any)
	if !ok {
		if nested, ok := contentMap["content"].(map[string]any); ok {
			items, ok = nested["richText"].([]any)
		}
	}
	if !ok || len(items) == 0 {
		return "", nil
	}

	parts := make([]string, 0, len(items))
	mediaPaths := make([]string, 0)
	for _, item := range items {
		block := normalizeMap(item)
		if len(block) == 0 {
			continue
		}

		blockText := firstNonEmpty(
			stringValue(block["text"]),
			stringValue(block["content"]),
			stringValue(block["title"]),
		)
		if blockText != "" {
			parts = append(parts, blockText)
		}

		downloadCode := firstNonEmpty(
			stringValue(block["pictureDownloadCode"]),
			stringValue(block["downloadCode"]),
		)
		if downloadCode == "" {
			continue
		}

		name := firstNonEmpty(
			stringValue(block["fileName"]),
			stringValue(block["name"]),
			"dingtalk-richtext-image.png",
		)
		path := c.downloadMessageFile(ctx, corpID, downloadCode, name)
		if path == "" {
			parts = append(parts, "[image]")
			continue
		}
		parts = append(parts, fmt.Sprintf("[image: %s]", filepath.Base(path)))
		mediaPaths = append(mediaPaths, path)
	}

	return strings.TrimSpace(strings.Join(parts, "\n")), mediaPaths
}

func (c *DingTalkChannel) handleDingTalkDownloadCode(
	ctx context.Context,
	corpID string,
	contentMap map[string]any,
	keys []string,
	baseName, ext, fallbackText string,
) (string, []string) {
	downloadCode := ""
	for _, key := range keys {
		downloadCode = firstNonEmpty(downloadCode, stringValue(contentMap[key]))
	}
	if downloadCode == "" {
		return fallbackText, nil
	}

	filename := firstNonEmpty(
		stringValue(contentMap["fileName"]),
		stringValue(contentMap["name"]),
		baseName+ext,
	)
	path := c.downloadMessageFile(ctx, corpID, downloadCode, filename)
	if path == "" {
		return fallbackText, nil
	}

	return fmt.Sprintf("%s %s", fallbackText, filepath.Base(path)), []string{path}
}

func (c *DingTalkChannel) downloadMessageFile(ctx context.Context, corpID, downloadCode, filename string) string {
	if downloadCode == "" {
		return ""
	}
	if corpID == "" {
		logger.WarnCF("dingtalk", "Skipping media download because corp ID is missing", map[string]any{
			"filename": filename,
		})
		return ""
	}

	accessToken, err := c.getCorpAccessToken(ctx, corpID)
	if err != nil {
		logger.ErrorCF("dingtalk", "Failed to get corp access token", map[string]any{
			"error":   err.Error(),
			"corp_id": corpID,
		})
		return ""
	}

	body, _ := json.Marshal(map[string]string{
		"downloadCode": downloadCode,
	})
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		"https://api.dingtalk.com/v1.0/robot/messageFiles/download",
		bytes.NewReader(body),
	)
	if err != nil {
		logger.ErrorCF("dingtalk", "Failed to create message file download request", map[string]any{
			"error": err.Error(),
		})
		return ""
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-acs-dingtalk-access-token", accessToken)

	resp, err := (&http.Client{Timeout: 30 * time.Second}).Do(req)
	if err != nil {
		logger.ErrorCF("dingtalk", "Failed to request message file download URL", map[string]any{
			"error": err.Error(),
		})
		return ""
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.ErrorCF("dingtalk", "Failed to read message file download response", map[string]any{
			"error": err.Error(),
		})
		return ""
	}
	if resp.StatusCode != http.StatusOK {
		logger.ErrorCF("dingtalk", "Message file download API returned non-200 status", map[string]any{
			"status": resp.StatusCode,
			"body":   string(respBody),
		})
		return ""
	}

	var result struct {
		DownloadURL string `json:"downloadUrl"`
	}
	if err := json.Unmarshal(respBody, &result); err != nil {
		logger.ErrorCF("dingtalk", "Failed to parse message file download response", map[string]any{
			"error": err.Error(),
			"body":  string(respBody),
		})
		return ""
	}
	if result.DownloadURL == "" {
		logger.WarnCF("dingtalk", "Message file download response did not contain downloadUrl", map[string]any{
			"body": string(respBody),
		})
		return ""
	}

	return utils.DownloadFile(result.DownloadURL, filename, utils.DownloadOptions{
		LoggerPrefix: "dingtalk",
	})
}

func (c *DingTalkChannel) getCorpAccessToken(ctx context.Context, corpID string) (string, error) {
	c.tokenMu.Lock()
	defer c.tokenMu.Unlock()

	if cached, ok := c.tokenCache[corpID]; ok && cached.Token != "" && time.Until(cached.ExpiresAt) > time.Minute {
		return cached.Token, nil
	}

	body, _ := json.Marshal(map[string]string{
		"suiteKey":    c.clientID,
		"suiteSecret": c.clientSecret,
		"authCorpId":  corpID,
	})
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		"https://api.dingtalk.com/v1.0/oauth2/corpAccessToken",
		bytes.NewReader(body),
	)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := (&http.Client{Timeout: 20 * time.Second}).Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("corpAccessToken returned %d: %s", resp.StatusCode, strings.TrimSpace(string(respBody)))
	}

	var result struct {
		AccessToken string `json:"accessToken"`
		ExpireIn    int    `json:"expireIn"`
	}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return "", err
	}
	if result.AccessToken == "" {
		return "", fmt.Errorf("corpAccessToken response missing accessToken")
	}

	expireIn := result.ExpireIn
	if expireIn <= 0 {
		expireIn = 7200
	}
	c.tokenCache[corpID] = dingtalkAccessToken{
		Token:     result.AccessToken,
		ExpiresAt: time.Now().Add(time.Duration(expireIn) * time.Second),
	}
	return result.AccessToken, nil
}

func structToMap(v any) map[string]any {
	raw, err := json.Marshal(v)
	if err != nil {
		return map[string]any{}
	}

	var decoded map[string]any
	if err := json.Unmarshal(raw, &decoded); err != nil {
		return map[string]any{}
	}
	return decoded
}

func normalizeMap(v any) map[string]any {
	switch value := v.(type) {
	case map[string]any:
		return value
	case nil:
		return map[string]any{}
	default:
		raw, err := json.Marshal(value)
		if err != nil {
			return map[string]any{}
		}
		var decoded map[string]any
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return map[string]any{}
		}
		return decoded
	}
}

func stringValue(v any) string {
	switch value := v.(type) {
	case string:
		return strings.TrimSpace(value)
	default:
		return ""
	}
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return strings.TrimSpace(value)
		}
	}
	return ""
}
