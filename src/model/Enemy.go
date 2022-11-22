package model

type Enemy struct {
	ID              string `gorm:"primaryKey"`
	SortID          int
	Name            string
	Index           string
	Race            string
	Level           string
	AttackType      string
	Endure          string
	Attack          string
	Defence         string
	Resistance      string
	Description     string
	Ability         string
	IsInvalidKilled bool
	EnemyInstances  []EnemyInstance
}
