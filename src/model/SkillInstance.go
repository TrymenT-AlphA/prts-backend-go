package model

type SkillInstance struct {
	Skill         Skill
	SkillID       string `gorm:"primarykey"`
	Level         int    `gorm:"primarykey"`
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