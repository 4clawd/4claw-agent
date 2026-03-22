package channels

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
)

type weChatAPIClient struct {
	baseURL string
	token   string
	client  *http.Client
}

func newWeChatAPIClient(baseURL, token string, timeout time.Duration) *weChatAPIClient {
	if timeout <= 0 {
		timeout = 15 * time.Second
	}
	return &weChatAPIClient{
		baseURL: strings.TrimRight(baseURL, "/"),
		token:   strings.TrimSpace(token),
		client:  &http.Client{Timeout: timeout},
	}
}

func (c *weChatAPIClient) postJSON(ctx context.Context, endpoint string, reqBody any, respBody any) error {
	raw, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/"+strings.TrimLeft(endpoint, "/"), bytes.NewReader(raw))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("AuthorizationType", "ilink_bot_token")
	req.Header.Set("X-WECHAT-UIN", randomWeChatUIN())
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("%s %d: %s", endpoint, resp.StatusCode, string(body))
	}
	if respBody == nil || len(body) == 0 {
		return nil
	}
	return json.Unmarshal(body, respBody)
}

func (c *weChatAPIClient) getJSON(ctx context.Context, endpoint string, params url.Values, headers map[string]string, respBody any) error {
	u, err := url.Parse(c.baseURL + "/" + strings.TrimLeft(endpoint, "/"))
	if err != nil {
		return err
	}
	if len(params) > 0 {
		u.RawQuery = params.Encode()
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("%s %d: %s", endpoint, resp.StatusCode, string(body))
	}
	if respBody == nil || len(body) == 0 {
		return nil
	}
	return json.Unmarshal(body, respBody)
}

func (c *weChatAPIClient) getUpdates(ctx context.Context, cursor string, timeoutMS int) (*weChatGetUpdatesResponse, error) {
	if timeoutMS <= 0 {
		timeoutMS = 35000
	}
	longPollClient := *c.client
	longPollClient.Timeout = time.Duration(timeoutMS+5000) * time.Millisecond
	longPoll := *c
	longPoll.client = &longPollClient
	var resp weChatGetUpdatesResponse
	err := longPoll.postJSON(ctx, "ilink/bot/getupdates", weChatGetUpdatesRequest{
		GetUpdatesBuf: cursor,
		BaseInfo:      weChatBaseInfo{ChannelVersion: "4claw-native"},
	}, &resp)
	return &resp, err
}

func (c *weChatAPIClient) sendMessage(ctx context.Context, message weChatMessage) error {
	return c.postJSON(ctx, "ilink/bot/sendmessage", weChatSendMessageRequest{Message: message}, nil)
}

func (c *weChatAPIClient) getUploadURL(ctx context.Context, req weChatGetUploadURLRequest) (*weChatGetUploadURLResponse, error) {
	req.BaseInfo = weChatBaseInfo{ChannelVersion: "4claw-native"}
	var resp weChatGetUploadURLResponse
	if err := c.postJSON(ctx, "ilink/bot/getuploadurl", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *weChatAPIClient) getConfig(ctx context.Context, userID, contextToken string) (*weChatGetConfigResponse, error) {
	var resp weChatGetConfigResponse
	err := c.postJSON(ctx, "ilink/bot/getconfig", weChatGetConfigRequest{
		ILinkUserID: userID,
		ContextToken: contextToken,
		BaseInfo:    weChatBaseInfo{ChannelVersion: "4claw-native"},
	}, &resp)
	return &resp, err
}

func (c *weChatAPIClient) sendTyping(ctx context.Context, userID, ticket string, status int) error {
	return c.postJSON(ctx, "ilink/bot/sendtyping", weChatSendTypingRequest{
		ILinkUserID:  userID,
		TypingTicket: ticket,
		Status:       status,
		BaseInfo:     weChatBaseInfo{ChannelVersion: "4claw-native"},
	}, nil)
}

func weChatFetchQRCode(ctx context.Context, baseURL string) (*weChatQRCodeResponse, error) {
	client := newWeChatAPIClient(baseURL, "", 15*time.Second)
	var resp weChatQRCodeResponse
	err := client.getJSON(ctx, path.Join("ilink", "bot", "get_bot_qrcode"), url.Values{
		"bot_type": []string{defaultWeChatBotType},
	}, nil, &resp)
	return &resp, err
}

func weChatPollQRStatus(ctx context.Context, baseURL, qrCode string) (*weChatQRStatusResponse, error) {
	client := newWeChatAPIClient(baseURL, "", 40*time.Second)
	var resp weChatQRStatusResponse
	err := client.getJSON(ctx, path.Join("ilink", "bot", "get_qrcode_status"), url.Values{
		"qrcode": []string{qrCode},
	}, map[string]string{
		"iLink-App-ClientVersion": "1",
	}, &resp)
	return &resp, err
}

func randomWeChatUIN() string {
	var b [4]byte
	if _, err := rand.Read(b[:]); err != nil {
		now := uint32(time.Now().UnixNano())
		binary.BigEndian.PutUint32(b[:], now)
	}
	n := binary.BigEndian.Uint32(b[:])
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d", n)))
}
