package model

type CharacterInstance struct {
	Character        Character `json:"character"`
	CharacterId      string    `json:"characterId" gorm:"primaryKey"`
	Phase            int64     `json:"phase" gorm:"primaryKey"`
	Level            int64     `json:"level" gorm:"primaryKey"`
	RangeId          string    `json:"rangeId"`
	MaxHp            int64     `json:"maxHp"`
	Atk              int64     `json:"atk"`
	Def              int64     `json:"def"`
	MagicResistance  float64   `json:"magicResistance"`
	Cost             int64     `json:"cost"`
	BlockCnt         int64     `json:"blockCnt"`
	MoveSpeed        float64   `json:"moveSpeed"`
	AttackSpeed      float64   `json:"attackSpeed"`
	BaseAttackTime   float64   `json:"baseAttackTime"`
	RespawnTime      int64     `json:"respawnTime"`
	HpRecoveryPerSec float64   `json:"hpRecoveryPerSec"`
	SpRecoveryPerSec float64   `json:"spRecoveryPerSec"`
	MaxDeployCount   int64     `json:"maxDeployCount"`
	MaxDeckStackCnt  int64     `json:"maxDeckStackCnt"`
	TauntLevel       int64     `json:"tauntLevel"`
	MassLevel        int64     `json:"massLevel"`
	BaseForceLevel   int64     `json:"baseForceLevel"`
	StunImmune       bool      `json:"stunImmune"`
	SilenceImmune    bool      `json:"silenceImmune"`
	SleepImmune      bool      `json:"sleepImmune"`
	FrozenImmune     bool      `json:"frozenImmune"`
	LevitateImmune   bool      `json:"levitateImmune"`
	EvolveCost       string    `json:"evolveCost"`
}
