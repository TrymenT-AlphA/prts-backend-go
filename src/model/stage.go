package model

import (
	"database/sql"
	"time"
)

type Stage struct {
	Id                   string         `json:"id" gorm:"primaryKey"`
	Hard                 *Stage         `json:"hard" gorm:"foreignKey:HardId"`
	HardId               sql.NullString `json:"hardId" gorm:"unique"`
	Drops                []Drop         `json:"drops" gorm:"foreignKey:StageId"`
	EnemyInstances       []EI_S         `json:"enemyInstances" gorm:"foreignKey:StageId"`
	Type                 string         `json:"type"`
	Difficulty           string         `json:"difficulty"`
	PerformanceStageFlag string         `json:"performanceStageFlag"`
	DiffGroup            string         `json:"diffGroup"`
	UnlockCondition      string         `json:"unlockCondition"`
	LevelId              string         `json:"levelId"`
	ZoneId               string         `json:"zoneId"`
	Code                 string         `json:"code"`
	Name                 string         `json:"name"`
	Description          string         `json:"description"`
	DangerLevel          string         `json:"dangerLevel"`
	CanPractice          bool           `json:"canPractice"`
	PracticeTicketCost   int64          `json:"practiceTicketCost"`
	ApCost               int64          `json:"apCost"`
	Activity             string         `json:"activity"`
	ActivityName         string         `json:"activityName"`
	ActivityDisplayType  string         `json:"activityDisplayType"`
	StartTime            time.Time      `json:"startTime"`
	EndTime              time.Time      `json:"endTime"`
	CharacterLimit       int64          `json:"characterLimit"`
	MaxLifePoint         int64          `json:"maxLifePoint"`
	InitialCost          int64          `json:"initialCost"`
	MaxCost              int64          `json:"maxCost"`
}
