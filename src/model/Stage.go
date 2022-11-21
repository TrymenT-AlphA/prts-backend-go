package model

type Stage struct {
	ID                   string `gorm:"primarykey"`
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
	HardStageID          string `gorm:"index:unique"`
	Drops                []Drop
	EnemyInstances       []EI_S
}
