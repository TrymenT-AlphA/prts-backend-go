package model

type C_BS struct {
	Character       Character
	CharacterID     string `gorm:"primaryKey"`
	BuildingSkill   BuildingSkill
	BuildingSkillID string `gorm:"primaryKey"`
	Phase           int
	Level           int
}
