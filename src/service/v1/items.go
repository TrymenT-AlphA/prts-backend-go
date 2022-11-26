package service

import "prts-backend/src/model"

func Items() []model.Item {
	var result []model.Item
	db.
		Model(&model.Item{}).
		Find(&result)
	return result
}
