package telegram

import (
	"time"
)

// UserSession represents a Telegram user session
type UserSession struct {
	ChatID        int64
	UserID        int64
	Username      string
	FirstName     string
	LastName      string
	SessionID     string
	Model         string
	Messages      []Message
	LastActive    time.Time
	CreatedAt     time.Time
}

// Message represents a message in a session
type Message struct {
	ID        int
	Role      string // "user" or "assistant"
	Content   string
	Timestamp time.Time
}

// Command represents a Telegram command
type Command struct {
	Name        string
	Description string
	Handler     CommandHandler
}

// CommandHandler is a function that handles a command
type CommandHandler func(chatID int64, args []string) (string, error)

// TelegramConfig holds Telegram configuration
type TelegramConfig struct {
	BotToken   string
	AdminChatID int64
	Enabled     bool
}
