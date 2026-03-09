package providers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/4claw/4claw/pkg/auth"
)

// CodexCliAuth represents the ~/.codex/auth.json file structure.
type CodexCliAuth struct {
	AccessToken string `json:"access_token"`
	AccountID   string `json:"account_id"`
	IDToken     string `json:"id_token"`
	Tokens struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		AccountID    string `json:"account_id"`
		IDToken      string `json:"id_token"`
	} `json:"tokens"`
}

// ReadCodexCliCredentials reads OAuth tokens from the Codex CLI auth file.
// Expiry is estimated as file modification time + 1 hour, matching picoclaw-clone.
func ReadCodexCliCredentials() (accessToken, accountID string, expiresAt time.Time, err error) {
	authPath, err := resolveCodexAuthPath()
	if err != nil {
		return "", "", time.Time{}, err
	}

	data, err := os.ReadFile(authPath)
	if err != nil {
		return "", "", time.Time{}, fmt.Errorf("reading %s: %w", authPath, err)
	}

	var authFile CodexCliAuth
	if err = json.Unmarshal(data, &authFile); err != nil {
		return "", "", time.Time{}, fmt.Errorf("parsing %s: %w", authPath, err)
	}

	accessToken = authFile.Tokens.AccessToken
	if accessToken == "" {
		accessToken = authFile.AccessToken
	}
	if accessToken == "" {
		return "", "", time.Time{}, fmt.Errorf("no access_token in %s", authPath)
	}
	accountID = authFile.Tokens.AccountID
	if accountID == "" {
		accountID = authFile.AccountID
	}
	if accountID == "" {
		accountID = resolveCodexAccountID(authFile)
	}

	stat, err := os.Stat(authPath)
	if err != nil {
		expiresAt = time.Now().Add(time.Hour)
	} else {
		expiresAt = stat.ModTime().Add(time.Hour)
	}

	return accessToken, accountID, expiresAt, nil
}

// CreateCodexCliTokenSource reuses Codex CLI OAuth credentials from ~/.codex/auth.json.
func CreateCodexCliTokenSource() func() (string, string, error) {
	return func() (string, string, error) {
		token, accountID, expiresAt, err := ReadCodexCliCredentials()
		if err != nil {
			return "", "", fmt.Errorf("reading codex cli credentials: %w", err)
		}

		if time.Now().After(expiresAt) {
			return "", "", fmt.Errorf("codex cli credentials expired (auth.json last modified > 1h ago). Run: codex login")
		}

		return token, accountID, nil
	}
}

func resolveCodexAuthPath() (string, error) {
	codexHome := os.Getenv("CODEX_HOME")
	if codexHome == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("getting home dir: %w", err)
		}
		codexHome = filepath.Join(home, ".codex")
	}
	return filepath.Join(codexHome, "auth.json"), nil
}

func resolveCodexAccountID(authFile CodexCliAuth) string {
	if authFile.Tokens.AccountID != "" {
		return authFile.Tokens.AccountID
	}
	if authFile.AccountID != "" {
		return authFile.AccountID
	}
	if accountID := auth.ExtractAccountID(authFile.Tokens.IDToken); accountID != "" {
		return accountID
	}
	if accountID := auth.ExtractAccountID(authFile.Tokens.AccessToken); accountID != "" {
		return accountID
	}
	if accountID := auth.ExtractAccountID(authFile.IDToken); accountID != "" {
		return accountID
	}
	if accountID := auth.ExtractAccountID(authFile.AccessToken); accountID != "" {
		return accountID
	}
	return ""
}
