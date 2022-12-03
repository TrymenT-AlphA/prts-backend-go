package model

type EnemyInstance struct {
	EnemyId          string  `json:"enemyId" gorm:"primaryKey"`
	Level            int64   `json:"level" gorm:"primaryKey"`
	Stages           []EI_S  `json:"stages" gorm:"foreignKey:EnemyId,Level"`
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	MaxHp            int64   `json:"maxHp"`
	Atk              int64   `json:"atk"`
	Def              int64   `json:"def"`
	MagicResistance  float64 `json:"magicResistance"`
	Cost             int64   `json:"cost"`
	BlockCnt         int64   `json:"blockCnt"`
	MoveSpeed        float64 `json:"moveSpeed"`
	AttackSpeed      float64 `json:"attackSpeed"`
	BaseAttackTime   float64 `json:"baseAttackTime"`
	RespawnTime      float64 `json:"respawnTime"`
	HpRecoveryPerSec float64 `json:"HpRecoveryPerSec"`
	SpRecoveryPerSec float64 `json:"spRecoveryPerSec"`
	MaxDeployCount   int64   `json:"maxDeployCount"`
	MassLevel        int64   `json:"massLevel"`
	BaseForceLevel   int64   `json:"baseForceLevel"`
	TauntLevel       int64   `json:"tauntLevel"`
	SilenceImmune    bool    `json:"silenceImmune"`
	StunImmune       bool    `json:"stunImmune"`
	SleepImmune      bool    `json:"sleepImmune"`
	FrozenImmune     bool    `json:"frozenImmune"`
	LevitateImmune   bool    `json:"levitateImmune"`
	LifePointReduce  int64   `json:"lifePointReduce"`
	LevelType        int64   `json:"levelType"`
	RangeRadius      float64 `json:"rangeRadius"`
	NumOfExtraDrops  int64   `json:"numOfExtraDrops"`
}
