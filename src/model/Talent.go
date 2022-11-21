package model

type Talent struct {
	PrefabKey     int `gorm:"primarykey"`
	Phase         int `gorm:"primarykey"`
	Level         int `gorm:"primarykey"`
	PotentialRank int `gorm:"primarykey"`
	Name          string
	Description   string
	RangeID       string
	Character     Character
	CharacterID   string `gorm:"primarykey"`
}
