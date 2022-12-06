package model

type C_BS struct {
	Character       Character     `json:"character"`
	CharacterId     string        `json:"characterId" gorm:"primaryKey"`
	BuildingSkill   BuildingSkill `json:"buildingSkill"`
	BuildingSkillId string        `json:"buildingSkillId" gorm:"primaryKey"`
	UnlockPhase     int64         `json:"unlockPhase"`
	UnlockLevel     int64         `json:"unlockLevel"`
}
