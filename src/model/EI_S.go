package model

type EI_S struct {
	EnemyInstance      EnemyInstance
	EnemyID            string `gorm:"primarykey"`
	EnemyInstanceLevel int    `gorm:"primarykey"`
	Stage              Stage
	StageID            string `gorm:"primarykey"`
}