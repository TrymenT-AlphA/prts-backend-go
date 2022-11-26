package service

import "prts-backend/src/model"

func Item(id string) model.Item {
	var result model.Item
	db.
		Model(&model.Item{}).
		Where(&model.Item{ID: id}).
		Limit(1).
		Find(&result)
	return result
}
