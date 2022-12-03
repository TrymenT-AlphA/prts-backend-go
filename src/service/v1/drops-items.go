package apiv1

import (
	"prts-backend/src/model"
	"prts-backend/src/service"
)

func DropsItems() ([]model.Item, error) {
	var result []model.Item
	err := service.DB.
		Raw(`SELECT prts_item.*
			FROM prts_item
			JOIN prts_drop
			ON prts_item.ID = prts_drop.ItemID
			GROUP BY ID
			ORDER BY IconID`).
		Scan(&result).
		Error
	return result, err
}
