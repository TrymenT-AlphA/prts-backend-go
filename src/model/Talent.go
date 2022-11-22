package model

type Talent struct {
	PrefabKey     int `gorm:"primaryKey"`
	Phase         int `gorm:"primaryKey"`
	Level         int `gorm:"primaryKey"`
	PotentialRank int `gorm:"primaryKey"`
	Name          string
	Description   string
	RangeID       string
	Character     Character
	CharacterID   string `gorm:"primaryKey"`
}
