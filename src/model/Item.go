package model

type Item struct {
	Id             string `json:"id" gorm:"primaryKey"`
	Drops          []Drop `json:"drops" gorm:"foreignKey:ItemId"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Rarity         int64  `json:"rarity"`
	IconId         string `json:"iconId"`
	SortId         int64  `json:"sortId"`
	Usage          string `json:"usage"`
	ObtainApproach string `json:"obtainApproach"`
	ClassifyType   string `json:"classifyType"`
	Type           string `json:"type"`
}
