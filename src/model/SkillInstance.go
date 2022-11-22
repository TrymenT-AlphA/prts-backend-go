package model

type SkillInstance struct {
	Skill         Skill
	SkillID       string `gorm:"primaryKey"`
	Level         int    `gorm:"primaryKey"`
	Name          string
	RangeID       string
	Description   string
	Type          int
	DurationType  int
	SpType        int
	MaxChargeTime int
	SpCost        int
	InitSp        int
	Duration      int
}
