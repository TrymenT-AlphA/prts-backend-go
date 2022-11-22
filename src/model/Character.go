package model

type Character struct {
	ID                          string `gorm:"primaryKey"`
	Name                        string
	Description                 string
	CanUseGeneralPotentialItem  bool
	CanUseActivityPotentialItem bool
	PotentialItemID             string
	ActivityPotentialItemID     string
	NationID                    string
	GroupID                     string
	TeamID                      string
	DisplayNumber               string
	Appellation                 string
	Position                    string
	TagList                     string
	ItemUsage                   string
	ItemDesc                    string
	ItemObtainApproach          string
	IsNotObtainable             bool
	IsSpChar                    bool
	MaxPotentialLevel           int
	Rarity                      int
	Profession                  string
	SubProfessionID             string
	AllSkillLvlupList           string
	PotentialList               string
	Token                       *Character
	TokenID                     string `gorm:"index:unique"`
	KeyFrameList                []CharacterInstance
	BuildingSkillList           []C_BS
	SkillList                   []C_S
	TalentList                  []Talent
}
