package model

type CharacterInstance struct {
	Character        Character
	CharacterID      string `gorm:"primarykey"`
	Phase            int    `gorm:"primarykey"`
	Level            int    `gorm:"primarykey"`
	RangeID          string
	MaxHp            int
	Atk              int
	Def              int
	MagicResistance  float64
	Cost             int
	BlockCnt         int
	MoveSpeed        float64
	AttackSpeed      float64
	BaseAttackTime   float64
	RespawnTime      int
	HpRecoveryPerSec float64
	SpRecoveryPerSec float64
	MaxDeployCount   int
	MaxDeckStackCnt  int
	TauntLevel       int
	MassLevel        int
	BaseForceLevel   int
	StunImmune       bool
	SilenceImmune    bool
	SleepImmune      bool
	FrozenImmune     bool
	LevitateImmune   bool
	EvolveCost       string
}