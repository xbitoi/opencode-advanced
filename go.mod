module github.com/xbitoi/opencode-advanced

go 1.24

require (
	// Telegram Bot
	github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.1

	// UI and Terminal
	github.com/charmbracelet/bubbletea v0.26.1
	github.com/charmbracelet/lipgloss v0.10.0
	github.com/charmbracelet/log v0.4.0

	// AI and LLM
	github.com/google/generative-ai-go v0.14.0
	github.com/sashabaranov/go-openai v1.24.0

	// Browser Automation
	github.com/chromedp/chromedp v1.14.1
	github.com/chromedp/cdproto v0.0.0-20240417011944-fdf5b095e970

	// HTTP and Networking
	github.com/go-resty/resty/v2 v2.12.0
	github.com/cloudflare/cloudflare-go v0.97.0

	// Database
	github.com/mattn/go-sqlite3 v1.14.20
	gorm.io/gorm v1.25.7
	gorm.io/driver/sqlite v1.5.5

	// JSON and Data Processing
	github.com/tidwall/gjson v1.17.0
	github.com/tidwall/sjson v1.2.5

	// Config Management
	github.com/spf13/cobra v1.8.0
	github.com/spf13/viper v1.18.2

	// File Operations
	github.com/otiai10/copy v1.14.0

	// Encryption
	golang.org/x/crypto v0.24.0

	// HTTP Server
	github.com/gin-gonic/gin v1.10.0

	// Utils
	github.com/google/uuid v1.6.0
	github.com/jedib0t/go-pretty/v6 v6.5.9
	github.com/fatih/color v1.17.0
)
