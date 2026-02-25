// 4claw - Ultra-lightweight personal AI agent
// License: MIT

package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/4claw/4claw/pkg/config"
	"github.com/4claw/4claw/pkg/skills"
	"github.com/4claw/4claw/pkg/utils"
)

func skillsHelp() {
	fmt.Println("\nSkills commands:")
	fmt.Println("  list                    List installed skills")
	fmt.Println("  install <repo>          Install skill from GitHub")
	fmt.Println("  install-builtin         Install all builtin skills to workspace")
	fmt.Println("  list-builtin            List available builtin skills")
	fmt.Println("  remove <name>           Remove installed skill")
	fmt.Println("  search                  Search available skills")
	fmt.Println("  show <name>             Show skill details")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  4claw skills list")
	fmt.Println("  4claw skills install 4claw/4claw-skills/weather")
	fmt.Println("  4claw skills install-builtin")
	fmt.Println("  4claw skills list-builtin")
	fmt.Println("  4claw skills remove weather")
	fmt.Println("  4claw skills install --registry clawhub github")
}

func skillsListCmd(loader *skills.SkillsLoader) {
	allSkills := loader.ListSkills()
	if len(allSkills) == 0 {
		fmt.Println("No skills installed.")
		return
	}

	fmt.Println("\nInstalled Skills:")
	fmt.Println("------------------")
	for _, skill := range allSkills {
		fmt.Printf("  - %s (%s)\n", skill.Name, skill.Source)
		if skill.Description != "" {
			fmt.Printf("    %s\n", skill.Description)
		}
	}
}

func skillsInstallCmd(installer *skills.SkillInstaller, cfg *config.Config) {
	if len(os.Args) < 4 {
		fmt.Println("Usage: 4claw skills install <github-repo>")
		fmt.Println("       4claw skills install --registry <name> <slug>")
		return
	}

	if os.Args[3] == "--registry" {
		if len(os.Args) < 6 {
			fmt.Println("Usage: 4claw skills install --registry <name> <slug>")
			fmt.Println("Example: 4claw skills install --registry clawhub github")
			return
		}
		registryName := os.Args[4]
		slug := os.Args[5]
		skillsInstallFromRegistry(cfg, registryName, slug)
		return
	}

	repo := os.Args[3]
	fmt.Printf("Installing skill from %s...\n", repo)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := installer.InstallFromGitHub(ctx, repo); err != nil {
		fmt.Printf("[ERROR] Failed to install skill: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[OK] Skill '%s' installed successfully.\n", filepath.Base(repo))
}

func skillsInstallFromRegistry(cfg *config.Config, registryName, slug string) {
	err := utils.ValidateSkillIdentifier(registryName)
	if err != nil {
		fmt.Printf("[ERROR] Invalid registry name: %v\n", err)
		os.Exit(1)
	}

	err = utils.ValidateSkillIdentifier(slug)
	if err != nil {
		fmt.Printf("[ERROR] Invalid slug: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Installing skill '%s' from %s registry...\n", slug, registryName)

	registryMgr := skills.NewRegistryManagerFromConfig(skills.RegistryConfig{
		MaxConcurrentSearches: cfg.Tools.Skills.MaxConcurrentSearches,
		ClawHub:               skills.ClawHubConfig(cfg.Tools.Skills.Registries.ClawHub),
	})

	registry := registryMgr.GetRegistry(registryName)
	if registry == nil {
		fmt.Printf("[ERROR] Registry '%s' not found or not enabled. Check your config.json.\n", registryName)
		os.Exit(1)
	}

	workspace := cfg.WorkspacePath()
	targetDir := filepath.Join(workspace, "skills", slug)

	if _, err = os.Stat(targetDir); err == nil {
		fmt.Printf("[ERROR] Skill '%s' is already installed at %s\n", slug, targetDir)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if err = os.MkdirAll(filepath.Join(workspace, "skills"), 0o755); err != nil {
		fmt.Printf("[ERROR] Failed to create skills directory: %v\n", err)
		os.Exit(1)
	}

	result, err := registry.DownloadAndInstall(ctx, slug, "", targetDir)
	if err != nil {
		rmErr := os.RemoveAll(targetDir)
		if rmErr != nil {
			fmt.Printf("[ERROR] Failed to remove partial install: %v\n", rmErr)
		}
		fmt.Printf("[ERROR] Failed to install skill: %v\n", err)
		os.Exit(1)
	}

	if result.IsMalwareBlocked {
		rmErr := os.RemoveAll(targetDir)
		if rmErr != nil {
			fmt.Printf("[ERROR] Failed to remove partial install: %v\n", rmErr)
		}
		fmt.Printf("[ERROR] Skill '%s' is flagged as malicious and cannot be installed.\n", slug)
		os.Exit(1)
	}

	if result.IsSuspicious {
		fmt.Printf("[WARN] Skill '%s' is flagged as suspicious.\n", slug)
	}

	fmt.Printf("[OK] Skill '%s' v%s installed successfully.\n", slug, result.Version)
	if result.Summary != "" {
		fmt.Printf("  %s\n", result.Summary)
	}
}

func skillsRemoveCmd(installer *skills.SkillInstaller, skillName string) {
	fmt.Printf("Removing skill '%s'...\n", skillName)

	if err := installer.Uninstall(skillName); err != nil {
		fmt.Printf("[ERROR] Failed to remove skill: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[OK] Skill '%s' removed successfully.\n", skillName)
}

func skillsInstallBuiltinCmd(workspace string) {
	builtinSkillsDir := "./4claw/skills"
	workspaceSkillsDir := filepath.Join(workspace, "skills")

	fmt.Println("Copying builtin skills to workspace...")

	skillsToInstall := []string{
		"weather",
		"news",
		"stock",
		"calculator",
	}

	for _, skillName := range skillsToInstall {
		builtinPath := filepath.Join(builtinSkillsDir, skillName)
		workspacePath := filepath.Join(workspaceSkillsDir, skillName)

		if _, err := os.Stat(builtinPath); err != nil {
			fmt.Printf("[WARN] Builtin skill '%s' not found: %v\n", skillName, err)
			continue
		}

		if err := os.MkdirAll(workspacePath, 0o755); err != nil {
			fmt.Printf("[ERROR] Failed to create directory for %s: %v\n", skillName, err)
			continue
		}

		if err := copyDirectory(builtinPath, workspacePath); err != nil {
			fmt.Printf("[ERROR] Failed to copy %s: %v\n", skillName, err)
		}
	}

	fmt.Println("\n[OK] Builtin skills installation finished.")
	fmt.Println("You can now use them in your workspace.")
}

func skillsListBuiltinCmd() {
	cfg, err := loadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}
	builtinSkillsDir := filepath.Join(filepath.Dir(cfg.WorkspacePath()), "4claw", "skills")

	fmt.Println("\nAvailable Builtin Skills:")
	fmt.Println("-----------------------")

	entries, err := os.ReadDir(builtinSkillsDir)
	if err != nil {
		fmt.Printf("Error reading builtin skills: %v\n", err)
		return
	}

	if len(entries) == 0 {
		fmt.Println("No builtin skills available.")
		return
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		skillName := entry.Name()
		skillFile := filepath.Join(builtinSkillsDir, skillName, "SKILL.md")
		description := "No description"

		if data, readErr := os.ReadFile(skillFile); readErr == nil {
			for _, line := range strings.Split(string(data), "\n") {
				trimmed := strings.TrimSpace(line)
				if strings.HasPrefix(strings.ToLower(trimmed), "description:") {
					description = strings.TrimSpace(strings.TrimPrefix(trimmed, "description:"))
					break
				}
			}
		}

		fmt.Printf("  [OK] %s\n", skillName)
		if description != "" {
			fmt.Printf("       %s\n", description)
		}
	}
}

func skillsSearchCmd(installer *skills.SkillInstaller) {
	fmt.Println("Searching for available skills...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	availableSkills, err := installer.ListAvailableSkills(ctx)
	if err != nil {
		fmt.Printf("[ERROR] Failed to fetch skills list: %v\n", err)
		return
	}

	if len(availableSkills) == 0 {
		fmt.Println("No skills available.")
		return
	}

	fmt.Printf("\nAvailable Skills (%d):\n", len(availableSkills))
	fmt.Println("--------------------")
	for _, skill := range availableSkills {
		fmt.Printf("  - %s\n", skill.Name)
		fmt.Printf("    %s\n", skill.Description)
		fmt.Printf("    Repo: %s\n", skill.Repository)
		if len(skill.Tags) > 0 {
			fmt.Printf("    Tags: %v\n", skill.Tags)
		}
		fmt.Println()
	}
}

func skillsShowCmd(loader *skills.SkillsLoader, skillName string) {
	content, ok := loader.LoadSkill(skillName)
	if !ok {
		fmt.Printf("[ERROR] Skill '%s' not found\n", skillName)
		return
	}

	fmt.Printf("\nSkill: %s\n", skillName)
	fmt.Println("----------------------")
	fmt.Println(content)
}
