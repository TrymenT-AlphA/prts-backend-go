package model

import (
	"database/sql"
	"time"
)

type Drop struct {
	Item     Item
	ItemID   string `gorm:"primaryKey"`
	Stage    Stage
	StageID  string `gorm:"primaryKey"`
	Times    int
	Quantity int
	StdDev   float64
	Start    time.Time `gorm:"primaryKey"`
	End      sql.NullTime
	UpdateAt time.Time
}
