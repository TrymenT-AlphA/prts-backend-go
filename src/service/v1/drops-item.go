package service

import (
	"prts-backend/src/model"
)

func DropsItem(id string) model.Item {
	var result model.Item
	db.
		Model(&model.Item{}).
		Where(&model.Item{ID: id}).
		Limit(1).
		Preload("Drops").
		Preload("Drops.Item").
		Preload("Drops.Stage").
		Find(&result)
	return result
}
