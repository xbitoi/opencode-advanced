package models

import (
	"fmt"
	"sync"
)

// ModelType represents a type of AI model
type ModelType string

const (
	Gemini    ModelType = "gemini"
	Claude    ModelType = "claude"
	GPT       ModelType = "gpt"
	DeepSeek  ModelType = "deepseek"
)

// ModelInfo holds information about an AI model
type ModelInfo struct {
	ID          string
	Name        string
	Provider    string
	Type        ModelType
	MaxTokens   int
	ContextSize int
	Version     string
	Enabled     bool
}

// Selector manages model selection
type Selector struct {
	models  map[string]*ModelInfo
	default string
	mutex   sync.RWMutex
}

// NewSelector creates a new model selector
func NewSelector() *Selector {
	return &Selector{
		models: make(map[string]*ModelInfo),
	}
}

// RegisterModel registers a model
func (s *Selector) RegisterModel(model *ModelInfo) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.models[model.ID] = model
	if s.default == "" {
		s.default = model.ID
	}
	return nil
}

// SelectModel selects a model by ID
func (s *Selector) SelectModel(modelID string) (*ModelInfo, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	model, exists := s.models[modelID]
	if !exists {
		return nil, fmt.Errorf("model %s not found", modelID)
	}

	return model, nil
}

// GetDefault returns the default model
func (s *Selector) GetDefault() (*ModelInfo, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if s.default == "" {
		return nil, fmt.Errorf("no default model set")
	}

	model, exists := s.models[s.default]
	if !exists {
		return nil, fmt.Errorf("default model not found")
	}

	return model, nil
}

// ListModels lists all available models
func (s *Selector) ListModels() []*ModelInfo {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var models []*ModelInfo
	for _, model := range s.models {
		if model.Enabled {
			models = append(models, model)
		}
	}
	return models
}
