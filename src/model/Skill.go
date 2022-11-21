package model

type Skill struct {
	ID             string `gorm:"primarykey"`
	IconID         string
	SkillInstances []SkillInstance
	Characters     []C_S
}