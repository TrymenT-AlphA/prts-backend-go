package model

type Talent struct {
	Character     Character
	CharacterID   string `gorm:"primaryKey"`
	PotentialRank int    `gorm:"primaryKey"`
	PrefabKey     int    `gorm:"primaryKey"`
	Phase         int    `gorm:"primaryKey"`
	Level         int    `gorm:"primaryKey"`
	Name          string
	Description   string
	RangeID       string
}
