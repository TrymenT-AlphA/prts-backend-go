package service

import "prts-backend/src/model"

func DropsItems() []model.Item {
	var result []model.Item
	db.
		Model(&model.Item{}).
		Having("COUNT(Drops) >= 1").
		Find(&result)
	return result
}
