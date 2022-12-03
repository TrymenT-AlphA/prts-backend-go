package model

type SkillInstance struct {
	Skill         Skill  `json:"skill"`
	SkillId       string `json:"skillId" gorm:"primaryKey"`
	Level         int64  `json:"level" gorm:"primaryKey"`
	Name          string `json:"name"`
	RangeId       string `json:"rangeId"`
	Description   string `json:"description"`
	Type          int64  `json:"type"`
	DurationType  int64  `json:"durationType"`
	SpType        int64  `json:"spType"`
	MaxChargeTime int64  `json:"maxChargeTime"`
	SpCost        int64  `json:"spCost"`
	InitSp        int64  `json:"initSp"`
	Duration      int64  `json:"duration"`
}
