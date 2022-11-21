package model

import "time"

type Drop struct {
	Item     Item
	ItemID   string `gorm:"primarykey"`
	Stage    Stage
	StageID  string `gorm:"primarykey"`
	Times    int
	Quantity int
	StdDev   float64
	Start    time.Time
	End      time.Time
	UpdateAt time.Time
}
