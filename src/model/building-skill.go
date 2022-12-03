package model

type BuildingSkill struct {
	Id          string `json:"id" gorm:"primaryKey"`
	Characters  []C_BS `json:"characters" gorm:"foreignKey:BuildingSkillId"`
	Icon        string `json:"icon"`
	Name        string `json:"name"`
	SortId      int64    `json:"sortId"`
	Category    string `json:"category"`
	RoomType    string `json:"roomType"`
	Description string `json:"description"`
}
