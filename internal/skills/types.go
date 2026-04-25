package skills

import (
	"time"
)

// Skill represents an AI skill that can be learned and executed
type Skill struct {
	ID          string
	Name        string
	Description string
	Category    string
	Version     string
	Enabled     bool
	Prompt      string
	Examples    []string
	Requires    []string // Dependencies on other skills
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Metadata    map[string]interface{}
}

// SkillConfig holds skill configuration
type SkillConfig struct {
	Directory string
	AutoLoad  bool
	AutoSave  bool
}

// SkillManager manages skills
type SkillManager interface {
	// AddSkill adds a new skill
	AddSkill(skill *Skill) error

	// RemoveSkill removes a skill
	RemoveSkill(skillID string) error

	// GetSkill retrieves a skill by ID
	GetSkill(skillID string) (*Skill, error)

	// ListSkills lists all skills
	ListSkills() ([]*Skill, error)

	// ExecuteSkill executes a skill
	ExecuteSkill(skillID string, input string) (string, error)

	// UpdateSkill updates a skill
	UpdateSkill(skill *Skill) error
}
