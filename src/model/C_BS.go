package model

type C_BS struct {
	Character       Character
	CharacterID     string `gorm:"primarykey"`
	BuildingSkill   BuildingSkill
	BuildingSkillID string `gorm:"primarykey"`
	Phase           int
	Level           int
}