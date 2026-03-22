package channels

import (
	"context"
	"fmt"
	"time"

	"github.com/4claw/4claw/pkg/config"
	"github.com/4claw/4claw/pkg/logger"
)

type WeChatLoginResult struct {
	AccountID string
	BaseURL   string
	UserID    string
}

func LoginWeChatWithQR(ctx context.Context, cfg config.WeChatConfig) (*WeChatLoginResult, error) {
	baseURL := cfg.BaseURL
	if baseURL == "" {
		baseURL = defaultWeChatBaseURL
	}
	qrResp, err := weChatFetchQRCode(ctx, baseURL)
	if err != nil {
		return nil, err
	}
	if qrResp == nil || qrResp.QRCode == "" || qrResp.QRCodeImageURL == "" {
		return nil, fmt.Errorf("wechat login did not return a valid QR code")
	}

	fmt.Println("Open this WeChat QR login URL and scan it with WeChat:")
	fmt.Println()
	fmt.Println(qrResp.QRCodeImageURL)
	fmt.Println()
	fmt.Println("Waiting for confirmation...")

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		status, err := weChatPollQRStatus(ctx, baseURL, qrResp.QRCode)
		if err != nil {
			return nil, err
		}
		switch status.Status {
		case "wait":
		case "scaned":
			fmt.Println("QR scanned. Confirm the login in WeChat.")
		case "expired":
			return nil, fmt.Errorf("wechat QR code expired; restart login")
		case "confirmed":
			if status.ILinkBotID == "" || status.BotToken == "" {
				return nil, fmt.Errorf("wechat login confirmed but missing bot credentials")
			}
			accountID := status.ILinkBotID
			accountBaseURL := status.BaseURL
			if accountBaseURL == "" {
				accountBaseURL = baseURL
			}
			err := saveWeChatAccount(accountID, weChatAccountData{
				AccountID: accountID,
				Token:     status.BotToken,
				BaseURL:   accountBaseURL,
				UserID:    status.ILinkUserID,
				SavedAt:   time.Now().UTC().Format(time.RFC3339),
			})
			if err != nil {
				return nil, err
			}
			logger.InfoCF("wechat", "Saved WeChat account", map[string]any{
				"account_id": accountID,
				"user_id":    status.ILinkUserID,
			})
			return &WeChatLoginResult{
				AccountID: accountID,
				BaseURL:   accountBaseURL,
				UserID:    status.ILinkUserID,
			}, nil
		}

		select {
		case <-time.After(1 * time.Second):
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}

func ListWeChatAccounts() ([]WeChatLoginResult, error) {
	accounts, err := listWeChatAccounts()
	if err != nil {
		return nil, err
	}
	out := make([]WeChatLoginResult, 0, len(accounts))
	for _, account := range accounts {
		out = append(out, WeChatLoginResult{
			AccountID: account.AccountID,
			BaseURL:   account.BaseURL,
			UserID:    account.UserID,
		})
	}
	return out, nil
}
