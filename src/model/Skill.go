package model

type Skill struct {
	ID             string `gorm:"primaryKey"`
	IconID         string
	SkillInstances []SkillInstance `gorm:"foreignKey:SkillID"`
	Characters     []C_S
}
