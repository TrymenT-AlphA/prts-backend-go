package model

type Enemy struct {
	Id              string          `json:"id" gorm:"primaryKey"`
	EnemyInstances  []EnemyInstance `json:"enemyInstances" grom:"foreignKey:EnemyId"`
	SortId          int64           `json:"sortId"`
	Name            string          `json:"name"`
	Index           string          `json:"index"`
	Race            string          `json:"race"`
	Level           string          `json:"level"`
	AttackType      string          `json:"attackType"`
	Endure          string          `json:"endure"`
	Attack          string          `json:"attack"`
	Defence         string          `json:"defence"`
	Resistance      string          `json:"resistance"`
	Description     string          `json:"description"`
	Ability         string          `json:"ability"`
	IsInvalidKilled bool            `json:"isInvalidKilled"`
}
