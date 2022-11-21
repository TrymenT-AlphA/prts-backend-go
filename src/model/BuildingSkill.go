package model

type BuildingSkill struct {
	ID          string `gorm:"primarykey"`
	Name        string
	Icon        string
	SortID      int
	Category    string
	RoomType    string
	Description string
	Characters  []C_BS
}