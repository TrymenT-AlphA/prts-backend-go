package model

type EI_S struct {
	EnemyInstance      EnemyInstance
	EnemyID            string `gorm:"primaryKey"`
	EnemyInstanceLevel int    `gorm:"primaryKey"`
	Stage              Stage
	StageID            string `gorm:"primaryKey"`
}
