package model

type C_S struct {
	Character     Character
	CharacterID   string `gorm:"primarykey"`
	Skill         Skill
	SkillID       string `gorm:"primarykey"`
	LvlupCostCond string
}