package model

type Talent struct {
	Character     Character `gorm:"foreignLey:Id"`
	CharacterId   string    `gorm:"primaryKey" json:"-"`
	TalentId      int64     `gorm:"primaryKey" json:"prefabKey"`
	Phase         int64     `gorm:"primaryKey" json:"phase"`
	Level         int64     `gorm:"primaryKey" json:"level"`
	PotentialRank int64     `gorm:"primaryKey" json:"potentialRank"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	RangeId       string    `json:"rangeId"`
}
