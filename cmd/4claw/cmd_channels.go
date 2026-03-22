package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/4claw/4claw/pkg/channels"
	"github.com/4claw/4claw/pkg/config"
)

func channelsCmd() {
	if len(os.Args) < 3 {
		channelsHelp()
		return
	}

	switch os.Args[2] {
	case "wechat":
		weChatCommand()
	default:
		fmt.Printf("Unknown channels command: %s\n", os.Args[2])
		channelsHelp()
	}
}

func channelsHelp() {
	fmt.Println("\nChannels commands:")
	fmt.Println("  wechat login        Start WeChat QR login")
	fmt.Println("  wechat accounts     List stored WeChat accounts")
}

func weChatCommand() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: 4claw channels wechat <login|accounts> [--config <path>]")
		return
	}

	subcommand := os.Args[3]
	configPath := ""
	for i := 4; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "--config", "-c":
			if i+1 < len(os.Args) {
				configPath = os.Args[i+1]
				i++
			}
		}
	}

	switch subcommand {
	case "login":
		weChatLoginCmd(configPath)
	case "accounts":
		weChatAccountsCmd()
	default:
		fmt.Printf("Unknown WeChat subcommand: %s\n", subcommand)
	}
}

func weChatLoginCmd(configPath string) {
	cfg, err := loadConfigFromPath(configPath)
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Minute)
	defer cancel()

	account, err := channels.LoginWeChatWithQR(ctx, cfg.Channels.WeChat)
	if err != nil {
		fmt.Printf("WeChat login failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("WeChat login successful!")
	fmt.Printf("Account ID: %s\n", account.AccountID)
	if account.UserID != "" {
		fmt.Printf("User ID: %s\n", account.UserID)
	}

	if cfg.Channels.WeChat.AccountID == "" || strings.EqualFold(cfg.Channels.WeChat.AccountID, "default") {
		cfg.Channels.WeChat.AccountID = account.AccountID
		if cfg.Channels.WeChat.BaseURL == "" {
			cfg.Channels.WeChat.BaseURL = account.BaseURL
		}
		targetPath := getConfigPath()
		if configPath != "" {
			targetPath = configPath
		}
		if err := config.SaveConfig(targetPath, cfg); err != nil {
			fmt.Printf("Warning: failed to update config with WeChat account_id: %v\n", err)
		} else {
			fmt.Printf("Updated config with channels.wechat.account_id=%s\n", account.AccountID)
		}
	}
}

func weChatAccountsCmd() {
	accounts, err := channels.ListWeChatAccounts()
	if err != nil {
		fmt.Printf("Failed to list WeChat accounts: %v\n", err)
		os.Exit(1)
	}
	if len(accounts) == 0 {
		fmt.Println("No stored WeChat accounts.")
		return
	}
	fmt.Println("Stored WeChat accounts:")
	for _, account := range accounts {
		fmt.Printf("  %s", account.AccountID)
		if account.UserID != "" {
			fmt.Printf(" (user: %s)", account.UserID)
		}
		if account.BaseURL != "" {
			fmt.Printf(" base=%s", account.BaseURL)
		}
		fmt.Println()
	}
}
