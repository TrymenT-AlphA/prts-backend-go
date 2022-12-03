package model

type C_S struct {
	Character     Character `json:"character"`
	CharacterId   string    `json:"characterId" gorm:"primaryKey"`
	Skill         Skill     `json:"skill"`
	SkillId       string    `json:"skillId" gorm:"primaryKey"`
	LvlupCostCond string    `json:"lvlupCostCond"`
}
