package service

import "prts-backend/src/model"

func DropsItems() ([]model.Item, error) {
	var result []model.Item
	err := db.
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
