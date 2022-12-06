package model

import "database/sql"

type Character struct {
	Id                          string              `json:"id" gorm:"primaryKey"`
	Token                       *Character          `json:"token" gorm:"foreignKey:TokenId;references:Id"`
	TokenId                     sql.NullString      `json:"tokenId" gorm:"unique"`
	CharacterInstances          []CharacterInstance `json:"characterInstances" gorm:"foreignKey:CharacterId"`
	BuildingSkills              []C_BS              `json:"buildingSkills" gorm:"foreignKey:CharacterId"`
	Skills                      []C_S               `json:"skills" gorm:"foreignKey:CharacterId"`
	Talents                     []Talent            `json:"talents" gorm:"foreignKey:CharacterId"`
	Name                        string              `json:"name"`
	Description                 string              `json:"description"`
	CanUseGeneralPotentialItem  bool                `json:"canUseGeneralPotentialItem"`
	CanUseActivityPotentialItem bool                `json:"canUseActivityPotentialItem"`
	PotentialItemId             string              `json:"potentialItemId"`
	ActivityPotentialItemId     string              `json:"activityPotentialItemId"`
	NationId                    string              `json:"nationId"`
	GroupId                     string              `json:"groupId"`
	TeamId                      string              `json:"teamId"`
	DisplayNumber               string              `json:"displayNumber"`
	Appellation                 string              `json:"appellation"`
	Position                    string              `json:"position"`
	TagList                     string              `json:"tagList"`
	ItemUsage                   string              `json:"itemUsage"`
	ItemDesc                    string              `json:"itemDesc"`
	ItemObtainApproach          string              `json:"itemObtainApproach"`
	IsNotObtainable             bool                `json:"isNotObtainable"`
	IsSpChar                    bool                `json:"isSpChar"`
	MaxPotentialLevel           int64               `json:"maxPotentialLevel"`
	Rarity                      int64               `json:"rarity"`
	Profession                  string              `json:"profession"`
	SubProfessionId             string              `json:"subProfessionId"`
	AllSkillLvlupList           string              `json:"allSkillLvlupList"`
	PotentialList               string              `json:"potentialList"`
	Sex                         string              `json:"sex"`
}
