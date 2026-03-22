package channels

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/4claw/4claw/pkg/config"
	"github.com/4claw/4claw/pkg/logger"
	"github.com/4claw/4claw/pkg/routing"
)

const (
	defaultWeChatBaseURL    = "https://ilinkai.weixin.qq.com"
	defaultWeChatCDNBaseURL = "https://novac2c.cdn.weixin.qq.com/c2c"
	defaultWeChatBotType    = "3"
)

type weChatAccountData struct {
	AccountID string `json:"account_id,omitempty"`
	Token     string `json:"token,omitempty"`
	BaseURL   string `json:"base_url,omitempty"`
	UserID    string `json:"user_id,omitempty"`
	SavedAt   string `json:"saved_at,omitempty"`
}

type weChatSyncBufData struct {
	GetUpdatesBuf string `json:"get_updates_buf,omitempty"`
}

func weChatHomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return "."
	}
	return filepath.Join(home, ".4claw", "wechat")
}

func weChatAccountsDir() string {
	return filepath.Join(weChatHomeDir(), "accounts")
}

func weChatAccountPath(accountID string) string {
	return filepath.Join(weChatAccountsDir(), routing.NormalizeAccountID(accountID)+".json")
}

func weChatSyncBufPath(accountID string) string {
	return filepath.Join(weChatAccountsDir(), routing.NormalizeAccountID(accountID)+".sync.json")
}

func saveWeChatAccount(accountID string, data weChatAccountData) error {
	accountID = routing.NormalizeAccountID(accountID)
	if err := os.MkdirAll(weChatAccountsDir(), 0o755); err != nil {
		return err
	}
	data.AccountID = accountID
	raw, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(weChatAccountPath(accountID), raw, 0o600)
}

func loadWeChatAccount(accountID string) (*weChatAccountData, error) {
	accountID = routing.NormalizeAccountID(accountID)
	raw, err := os.ReadFile(weChatAccountPath(accountID))
	if err != nil {
		return nil, err
	}
	var data weChatAccountData
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, err
	}
	if data.AccountID == "" {
		data.AccountID = accountID
	}
	return &data, nil
}

func listWeChatAccounts() ([]weChatAccountData, error) {
	entries, err := os.ReadDir(weChatAccountsDir())
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	accounts := make([]weChatAccountData, 0, len(entries))
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || filepath.Ext(name) != ".json" || strings.HasSuffix(name, ".sync.json") {
			continue
		}
		data, err := loadWeChatAccount(strings.TrimSuffix(name, ".json"))
		if err != nil {
			logger.WarnCF("wechat", "Skipping unreadable WeChat account", map[string]any{
				"file":  name,
				"error": err.Error(),
			})
			continue
		}
		accounts = append(accounts, *data)
	}
	return accounts, nil
}

func saveWeChatSyncBuf(accountID, buf string) error {
	accountID = routing.NormalizeAccountID(accountID)
	if err := os.MkdirAll(weChatAccountsDir(), 0o755); err != nil {
		return err
	}
	raw, err := json.Marshal(weChatSyncBufData{GetUpdatesBuf: buf})
	if err != nil {
		return err
	}
	return os.WriteFile(weChatSyncBufPath(accountID), raw, 0o600)
}

func loadWeChatSyncBuf(accountID string) string {
	accountID = routing.NormalizeAccountID(accountID)
	raw, err := os.ReadFile(weChatSyncBufPath(accountID))
	if err != nil {
		return ""
	}
	var data weChatSyncBufData
	if err := json.Unmarshal(raw, &data); err != nil {
		return ""
	}
	return data.GetUpdatesBuf
}

type resolvedWeChatAccount struct {
	AccountID  string
	BaseURL    string
	CDNBaseURL string
	Token      string
	UserID     string
}

func resolveWeChatAccount(cfg config.WeChatConfig) (*resolvedWeChatAccount, error) {
	accountID := routing.NormalizeAccountID(cfg.AccountID)
	baseURL := cfg.BaseURL
	if baseURL == "" {
		baseURL = defaultWeChatBaseURL
	}
	cdnBaseURL := cfg.CDNBaseURL
	if cdnBaseURL == "" {
		cdnBaseURL = defaultWeChatCDNBaseURL
	}

	if cfg.Token != "" {
		return &resolvedWeChatAccount{
			AccountID:  accountID,
			BaseURL:    baseURL,
			CDNBaseURL: cdnBaseURL,
			Token:      cfg.Token,
		}, nil
	}

	if accountID == routing.DefaultAccountID {
		accounts, err := listWeChatAccounts()
		if err == nil {
			if len(accounts) == 1 {
				accountID = accounts[0].AccountID
			} else if len(accounts) > 1 {
				return nil, fmt.Errorf("multiple WeChat accounts are stored; set channels.wechat.account_id explicitly")
			}
		}
	}

	data, err := loadWeChatAccount(accountID)
	if err != nil {
		return nil, err
	}
	if data.BaseURL != "" {
		baseURL = data.BaseURL
	}
	return &resolvedWeChatAccount{
		AccountID:  accountID,
		BaseURL:    baseURL,
		CDNBaseURL: cdnBaseURL,
		Token:      data.Token,
		UserID:     data.UserID,
	}, nil
}
