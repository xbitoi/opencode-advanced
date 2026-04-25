package bridges

import (
	"context"
	"fmt"
)

// DeepSeekBridge implements the Bridge interface for DeepSeek
type DeepSeekBridge struct {
	config BridgeConfig
	model  string
}

// NewDeepSeekBridge creates a new DeepSeek bridge
func NewDeepSeekBridge() *DeepSeekBridge {
	return &DeepSeekBridge{
		model: "deepseek-chat",
	}
}

// Initialize initializes the DeepSeek bridge
func (d *DeepSeekBridge) Initialize(config BridgeConfig) error {
	if config.APIKey == "" && !config.BrowserMode {
		return fmt.Errorf("deepseek API key or browser mode is required")
	}

	d.config = config
	return nil
}

// Send sends a message to DeepSeek
func (d *DeepSeekBridge) Send(ctx context.Context, message string, sessionID string) (*BridgeResponse, error) {
	if !d.IsAvailable() {
		return nil, fmt.Errorf("deepseek bridge not available")
	}

	// TODO: Implement DeepSeek API or browser-based integration
	// For now, return a placeholder
	return &BridgeResponse{
		Content: "DeepSeek response placeholder - implement integration",
		Model:   d.model,
		Tokens:  0,
	}, nil
}

// IsAvailable checks if DeepSeek bridge is available
func (d *DeepSeekBridge) IsAvailable() bool {
	return d.config.Enabled && (d.config.APIKey != "" || d.config.BrowserMode)
}

// GetName returns the bridge name
func (d *DeepSeekBridge) GetName() string {
	return "DeepSeek"
}

// GetType returns the bridge type
func (d *DeepSeekBridge) GetType() BridgeType {
	return DeepSeekBridge
}

// Cleanup cleans up resources
func (d *DeepSeekBridge) Cleanup() error {
	return nil
}
