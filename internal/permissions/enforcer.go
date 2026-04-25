package permissions

import (
	"fmt"
	"sync"
)

// Permission represents a permission type
type Permission string

const (
	PermissionRead   Permission = "read"
	PermissionWrite  Permission = "write"
	PermissionDelete Permission = "delete"
	PermissionExecute Permission = "execute"
	PermissionModify Permission = "modify"
)

// Enforcer enforces permission rules
type Enforcer struct {
	rules      map[string]bool // permission:action -> allowed
	mutex      sync.RWMutex
	auditLog   []AuditEntry
	secureMode bool
}

// AuditEntry represents an audit log entry
type AuditEntry struct {
	Action     string
	Permission Permission
	Allowed    bool
	Timestamp  string
	Details    string
}

// NewEnforcer creates a new permission enforcer
func NewEnforcer(secureMode bool) *Enforcer {
	return &Enforcer{
		rules:      make(map[string]bool),
		secureMode: secureMode,
	}
}

// AllowPermission allows a permission for an action
func (e *Enforcer) AllowPermission(action string, permission Permission) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	key := fmt.Sprintf("%s:%s", action, permission)
	e.rules[key] = true
}

// DenyPermission denies a permission for an action
func (e *Enforcer) DenyPermission(action string, permission Permission) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	key := fmt.Sprintf("%s:%s", action, permission)
	e.rules[key] = false
}

// IsAllowed checks if a permission is allowed for an action
func (e *Enforcer) IsAllowed(action string, permission Permission) bool {
	e.mutex.RLock()
	defer e.mutex.RUnlock()

	if e.secureMode {
		// In secure mode, only explicitly allowed permissions are allowed
		key := fmt.Sprintf("%s:%s", action, permission)
		return e.rules[key]
	}

	// In normal mode, deny only explicitly denied permissions
	key := fmt.Sprintf("%s:%s", action, permission)
	allowed, exists := e.rules[key]
	if !exists {
		return true
	}
	return allowed
}

// AuditLog logs an action for audit purposes
func (e *Enforcer) AuditLog(action string, permission Permission, allowed bool, details string) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	entry := AuditEntry{
		Action:     action,
		Permission: permission,
		Allowed:    allowed,
		Details:    details,
	}
	e.auditLog = append(e.auditLog, entry)
}

// GetAuditLog returns the audit log
func (e *Enforcer) GetAuditLog() []AuditEntry {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return append([]AuditEntry{}, e.auditLog...)
}
