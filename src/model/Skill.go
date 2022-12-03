package model

type Skill struct {
	Id             string          `json:"id" gorm:"primaryKey"`
	SkillInstances []SkillInstance `json:"skillInstances" gorm:"foreignKey:SkillId"`
	Characters     []C_S           `json:"characters" gorm:"foreignKey:SkillId"`
	IconId         string          `json:"iconId"`
}
