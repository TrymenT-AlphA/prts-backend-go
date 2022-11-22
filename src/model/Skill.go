package model

type Skill struct {
	ID             string `gorm:"primaryKey"`
	IconID         string
	SkillInstances []SkillInstance
	Characters     []C_S
}
