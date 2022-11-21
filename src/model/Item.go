package model

type Item struct {
	ID             string `gorm:"primarykey"`
	Name           string
	Description    string
	Rarity         int
	IconID         string
	SortID         int
	Usage          string
	ObtainApproach string
	ClassifyType   string
	Type           string
	Drops          []Drop
}