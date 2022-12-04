package apiv1

import (
	"prts-backend/src/model"
	"prts-backend/src/service"
)

func DropsItems() ([]model.Item, error) {
	var result []model.Item
	err := service.DB.
		Raw(`SELECT prts_Item.*
			FROM prts_Item
			JOIN prts_Drop
			ON prts_Item.Id = prts_Drop.ItemId
			GROUP BY Id
			ORDER BY IconId`).
		Scan(&result).
		Error
	return result, err
}
