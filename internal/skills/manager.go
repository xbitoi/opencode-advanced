package skills

import (
	"fmt"
	"sync"
	"time"
)

// DefaultSkillManager implements the SkillManager interface
type DefaultSkillManager struct {
	skills map[string]*Skill
	mutex  sync.RWMutex
	config SkillConfig
}

// NewSkillManager creates a new skill manager
func NewSkillManager(config SkillConfig) *DefaultSkillManager {
	return &DefaultSkillManager{
		skills: make(map[string]*Skill),
		config: config,
	}
}

// AddSkill adds a new skill
func (m *DefaultSkillManager) AddSkill(skill *Skill) error {
	if skill.ID == "" {
		return fmt.Errorf("skill ID is required")
	}

	m.mutex.Lock()
	defer m.mutex.Unlock()

	if _, exists := m.skills[skill.ID]; exists {
		return fmt.Errorf("skill %s already exists", skill.ID)
	}

	skill.CreatedAt = time.Now()
	skill.UpdatedAt = time.Now()
	m.skills[skill.ID] = skill

	return nil
}

// RemoveSkill removes a skill
func (m *DefaultSkillManager) RemoveSkill(skillID string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if _, exists := m.skills[skillID]; !exists {
		return fmt.Errorf("skill %s not found", skillID)
	}

	delete(m.skills, skillID)
	return nil
}

// GetSkill retrieves a skill by ID
func (m *DefaultSkillManager) GetSkill(skillID string) (*Skill, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	skill, exists := m.skills[skillID]
	if !exists {
		return nil, fmt.Errorf("skill %s not found", skillID)
	}

	return skill, nil
}

// ListSkills lists all skills
func (m *DefaultSkillManager) ListSkills() ([]*Skill, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	var skills []*Skill
	for _, skill := range m.skills {
		if skill.Enabled {
			skills = append(skills, skill)
		}
	}

	return skills, nil
}

// ExecuteSkill executes a skill
func (m *DefaultSkillManager) ExecuteSkill(skillID string, input string) (string, error) {
	skill, err := m.GetSkill(skillID)
	if err != nil {
		return "", err
	}

	if !skill.Enabled {
		return "", fmt.Errorf("skill %s is disabled", skillID)
	}

	// TODO: Implement skill execution with prompt and input
	return fmt.Sprintf("Executing skill %s with input: %s", skillID, input), nil
}

// UpdateSkill updates a skill
func (m *DefaultSkillManager) UpdateSkill(skill *Skill) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if _, exists := m.skills[skill.ID]; !exists {
		return fmt.Errorf("skill %s not found", skill.ID)
	}

	skill.UpdatedAt = time.Now()
	m.skills[skill.ID] = skill

	return nil
}
