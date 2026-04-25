package main

import (
	"fmt"
	"os"

	"github.com/xbitoi/opencode-advanced/internal/bridges
	"github.com/xbitoi/opencode-advanced/internal/skills
	"github.com/xbitoi/opencode-advanced/internal/telegram
)

func main() {
	// Initialize skill manager
	skillConfig := skills.SkillConfig{
		Directory: "./.opencode/skills",
		AutoLoad:  true,
		AutoSave:  true,
	}
	skillMgr := skills.NewSkillManager(skillConfig)

	// Initialize bridge manager
	bridgeMgr := bridges.NewManager()

	// Register Gemini bridge
	geminiBridge := bridges.NewGeminiBridge()
	bridgeMgr.RegisterBridge(bridges.GeminiBridge, geminiBridge)

	// Register DeepSeek bridge
	deepseekBridge := bridges.NewDeepSeekBridge()
	bridgeMgr.RegisterBridge(bridges.DeepSeekBridge, deepseekBridge)

	// Initialize Telegram client (optional)
	var telegramClient *telegram.Client
	if botToken := os.Getenv("TELEGRAM_BOT_TOKEN"); botToken != "" {
		cfg := telegram.TelegramConfig{
			BotToken:   botToken,
			AdminChatID: 0, // Set from config
			Enabled:    true,
		}
		var err error
		telegramClient, err = telegram.NewClient(cfg)
		if err != nil {
			fmt.Printf("Warning: Failed to initialize Telegram client: %v\n", err)
		}
	}

	fmt.Println("OpenCode Advanced initialized successfully!")
	fmt.Printf("Skills Manager: %v\n", skillMgr)
	fmt.Printf("Bridge Manager: %v\n", bridgeMgr)
	fmt.Printf("Telegram Client: %v\n", telegramClient)

	// TODO: Implement main CLI logic
	fmt.Println("\nRunning OpenCode Advanced...")
}
