package model

import (
	"database/sql"
	"time"
)

type Drop struct {
	Item     Item         `json:"item"`
	ItemId   string       `json:"itemId" gorm:"primaryKey"`
	Stage    Stage        `json:"stage"`
	StageId  string       `json:"stageId" gorm:"primaryKey"`
	Start    time.Time    `json:"start" gorm:"primaryKey"`
	End      sql.NullTime `json:"end"`
	Times    int64        `json:"times"`
	Quantity int64        `json:"quantity"`
	StdDev   float64      `json:"stdDev"`
	UpdateAt time.Time    `json:"updateAt"`
}
