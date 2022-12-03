package model

type EI_S struct {
	EnemyInstance EnemyInstance `json:"enemyInstances" gorm:"foreignKey:EnemyId,Level"`
	EnemyId       string        `json:"enemyId" gorm:"primaryKey"`
	Level         int64           `json:"level" gorm:"primaryKey"`
	Stage         Stage         `json:"stage" gorm:"foreignKey:StageId"`
	StageId       string        `json:"stageId" gorm:"primaryKey"`
}
