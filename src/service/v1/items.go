package apiv1

import (
	"prts-backend/src/model"
	"prts-backend/src/service"
)

func Items() []model.Item {
	var result []model.Item
	service.DB.
		Model(&model.Item{}).
		Find(&result)
	return result
}
