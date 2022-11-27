package model

type EnemyInstance struct {
	EnemyID         string `gorm:"primaryKey"`
	Level           int    `gorm:"primaryKey"`
	MaxHp           int
	Atk             int
	Def             int
	MagicResistance float64
	LifePointReduce int
	AttackSpeed     float64
	BaseAttackTime  float64
	RangeRadius     float64
	MoveSpeed       float64
	MassLevel       int
	RespawnTime     float64
	SilenceImmune   bool
	StunImmune      bool
	SleepImmune     bool
	FrozenImmune    bool
	LevitateImmune  bool
	Stages          []EI_S `gorm:"foreignKey:EnemyInstanceEnemyID,EnemyInstanceLevel"`
}
