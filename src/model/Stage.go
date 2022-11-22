package model

import "database/sql"

type Stage struct {
	ID                   string `gorm:"primaryKey"`
	Type                 string
	Difficulty           string
	PerformanceStageFlag string
	DiffGroup            string
	UnlockCondition      string
	LevelID              string
	ZoneID               string
	Code                 string
	Name                 string
	Description          string
	DangerLevel          string
	CanPractice          bool
	PracticeTicketCost   int
	ApCost               int
	HardStage            *Stage
	HardStageID          sql.NullString `gorm:"index:unique"`
	Drops                []Drop
	EnemyInstances       []EI_S
}
