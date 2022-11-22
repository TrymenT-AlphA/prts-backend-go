package model

type C_S struct {
	Character     Character
	CharacterID   string `gorm:"primaryKey"`
	Skill         Skill
	SkillID       string `gorm:"primaryKey"`
	LvlupCostCond string
}
