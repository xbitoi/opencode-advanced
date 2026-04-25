package bridges

import (
	"context"
)

// BridgeType represents the type of bridge
type BridgeType string

const (
	GeminiBridge   BridgeType = "gemini"
	DeepSeekBridge BridgeType = "deepseek"
)

// BridgeResponse represents a response from an AI bridge
type BridgeResponse struct {
	Content    string
	Model      string
	Tokens     int
	Error      error
	RawData    map[string]interface{}
}

// BridgeConfig holds configuration for a bridge
type BridgeConfig struct {
	APIKey          string
	Enabled         bool
	UseBrowser      bool
	Headless        bool
	BrowserMode     bool
	AntiDetect      bool
	Timeout         int
	MaxRetries      int
	UseProxy        bool
	ProxyURL        string
}

// Bridge interface defines methods for AI bridge implementations
type Bridge interface {
	// Send sends a message to the AI model
	Send(ctx context.Context, message string, sessionID string) (*BridgeResponse, error)

	// IsAvailable checks if the bridge is available and configured
	IsAvailable() bool

	// GetName returns the name of the bridge
	GetName() string

	// GetType returns the type of the bridge
	GetType() BridgeType

	// Initialize initializes the bridge
	Initialize(config BridgeConfig) error

	// Cleanup cleans up resources
	Cleanup() error
}
