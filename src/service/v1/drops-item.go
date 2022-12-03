package apiv1

import (
	"prts-backend/src/model"
	"prts-backend/src/service"
)

func DropsItem(id string) model.Item {
	var result model.Item
	service.DB.
		Model(&model.Item{}).
		Where(&model.Item{Id: id}).
		Limit(1).
		Preload("Drops").
		Preload("Drops.Item").
		Preload("Drops.Stage").
		Find(&result)
	return result
}
