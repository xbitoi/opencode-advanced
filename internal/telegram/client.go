package telegram

import (
	"fmt"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Client represents a Telegram bot client
type Client struct {
	api      *tgbotapi.BotAPI
	config   TelegramConfig
	sessions map[int64]*UserSession
	mutex    sync.RWMutex
	commands map[string]CommandHandler
}

// NewClient creates a new Telegram client
func NewClient(config TelegramConfig) (*Client, error) {
	if config.BotToken == "" {
		return nil, fmt.Errorf("bot token is required")
	}

	api, err := tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		return nil, fmt.Errorf("failed to create bot API: %w", err)
	}

	return &Client{
		api:      api,
		config:   config,
		sessions: make(map[int64]*UserSession),
		commands: make(map[string]CommandHandler),
	}, nil
}

// RegisterCommand registers a command handler
func (c *Client) RegisterCommand(name string, handler CommandHandler) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.commands[name] = handler
}

// GetOrCreateSession gets or creates a user session
func (c *Client) GetOrCreateSession(chatID int64) *UserSession {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if session, exists := c.sessions[chatID]; exists {
		return session
	}

	session := &UserSession{
		ChatID: chatID,
	}
	c.sessions[chatID] = session
	return session
}

// SendMessage sends a message via Telegram
func (c *Client) SendMessage(chatID int64, message string) error {
	msg := tgbotapi.NewMessage(chatID, message)
	_, err := c.api.Send(msg)
	return err
}

// Start starts listening for updates
func (c *Client) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := c.api.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		// Handle message
		c.handleMessage(update.Message)
	}

	return nil
}

func (c *Client) handleMessage(message *tgbotapi.Message) {
	// TODO: Implement message handling
}
