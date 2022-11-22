package model

import "time"

type Drop struct {
	Item     Item
	ItemID   string `gorm:"primaryKey"`
	Stage    Stage
	StageID  string `gorm:"primaryKey"`
	Times    int
	Quantity int
	StdDev   float64
	Start    time.Time
	End      time.Time
	UpdateAt time.Time
}
