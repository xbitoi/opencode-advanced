package bridges

import (
	"context"
	"fmt"

	genai "google.golang.org/genai"
)

// GeminiBridge implements the Bridge interface for Google Gemini
type GeminiBridge struct {
	client *genai.Client
	config BridgeConfig
	model  string
}

// NewGeminiBridge creates a new Gemini bridge
func NewGeminiBridge() *GeminiBridge {
	return &GeminiBridge{
		model: "gemini-2.5-pro",
	}
}

// Initialize initializes the Gemini bridge
func (g *GeminiBridge) Initialize(config BridgeConfig) error {
	if config.APIKey == "" {
		return fmt.Errorf("gemini API key is required")
	}

	g.config = config

	// Initialize Gemini client
	client, err := genai.NewClient(&genai.ClientConfig{
		APIKey: config.APIKey,
	})
	if err != nil {
		return fmt.Errorf("failed to initialize gemini client: %w", err)
	}

	g.client = client
	return nil
}

// Send sends a message to Gemini
func (g *GeminiBridge) Send(ctx context.Context, message string, sessionID string) (*BridgeResponse, error) {
	if g.client == nil {
		return nil, fmt.Errorf("gemini bridge not initialized")
	}

	model := g.client.GenerativeModel(g.model)
	resp, err := model.GenerateContent(ctx, genai.Text(message))
	if err != nil {
		return nil, fmt.Errorf("failed to generate content: %w", err)
	}

	content := ""
	if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
		if text, ok := resp.Candidates[0].Content.Parts[0].(genai.Text); ok {
			content = string(text)
		}
	}

	return &BridgeResponse{
		Content: content,
		Model:   g.model,
		Tokens:  int(resp.UsageMetadata.InputTokenCount + resp.UsageMetadata.OutputTokenCount),
		RawData: map[string]interface{}{
			"usage": resp.UsageMetadata,
		},
	}, nil
}

// IsAvailable checks if Gemini bridge is available
func (g *GeminiBridge) IsAvailable() bool {
	return g.client != nil && g.config.Enabled
}

// GetName returns the bridge name
func (g *GeminiBridge) GetName() string {
	return "Google Gemini"
}

// GetType returns the bridge type
func (g *GeminiBridge) GetType() BridgeType {
	return GeminiBridge
}

// Cleanup cleans up resources
func (g *GeminiBridge) Cleanup() error {
	if g.client != nil {
		return g.client.Close()
	}
	return nil
}
