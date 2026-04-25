package bridges

import (
	"context"
	"fmt"
	"sync"
)

// Manager manages multiple AI bridges
type Manager struct {
	bridges map[BridgeType]Bridge
	mutex   sync.RWMutex
	default BridgeType
}

// NewManager creates a new bridge manager
func NewManager() *Manager {
	return &Manager{
		bridges: make(map[BridgeType]Bridge),
	}
}

// RegisterBridge registers a bridge
func (m *Manager) RegisterBridge(bridgeType BridgeType, bridge Bridge) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.bridges[bridgeType] = bridge
	if m.default == "" {
		m.default = bridgeType
	}
	return nil
}

// SetDefault sets the default bridge
func (m *Manager) SetDefault(bridgeType BridgeType) error {
	m.mutex.RLock()
	_, exists := m.bridges[bridgeType]
	m.mutex.RUnlock()

	if !exists {
		return fmt.Errorf("bridge %s not registered", bridgeType)
	}

	m.mutex.Lock()
	m.default = bridgeType
	m.mutex.Unlock()

	return nil
}

// GetBridge returns a bridge by type
func (m *Manager) GetBridge(bridgeType BridgeType) (Bridge, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	bridge, exists := m.bridges[bridgeType]
	if !exists {
		return nil, fmt.Errorf("bridge %s not found", bridgeType)
	}

	return bridge, nil
}

// GetDefault returns the default bridge
func (m *Manager) GetDefault() (Bridge, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	if m.default == "" {
		return nil, fmt.Errorf("no default bridge set")
	}

	bridge, exists := m.bridges[m.default]
	if !exists {
		return nil, fmt.Errorf("default bridge %s not found", m.default)
	}

	return bridge, nil
}

// ListAvailableBridges returns a list of available bridges
func (m *Manager) ListAvailableBridges() []Bridge {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	var available []Bridge
	for _, bridge := range m.bridges {
		if bridge.IsAvailable() {
			available = append(available, bridge)
		}
	}
	return available
}

// Send sends a message using the specified bridge or default
func (m *Manager) Send(ctx context.Context, message string, sessionID string, bridgeType *BridgeType) (*BridgeResponse, error) {
	var bridge Bridge
	var err error

	if bridgeType != nil {
		bridge, err = m.GetBridge(*bridgeType)
	} else {
		bridge, err = m.GetDefault()
	}

	if err != nil {
		return nil, err
	}

	return bridge.Send(ctx, message, sessionID)
}

// Cleanup cleans up all bridges
func (m *Manager) Cleanup() error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	var errors []error
	for _, bridge := range m.bridges {
		if err := bridge.Cleanup(); err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("cleanup errors: %v", errors)
	}
	return nil
}
