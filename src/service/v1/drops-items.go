package service

import "prts-backend/src/model"

func DropsItems() []model.Item {
	var result []model.Item
	db.
		Raw(`SELECT prts_item.*
			FROM prts_item
			JOIN prts_drop
			ON prts_item.ID = prts_drop.ItemID
			GROUP BY ID`).
		Scan(&result)
	return result
}
