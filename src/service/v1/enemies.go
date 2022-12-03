package apiv1

import (
	"prts-backend/src/model"
	"prts-backend/src/service"
)

func Enemies() []model.Enemy {
	var result []model.Enemy
	service.DB.
		Model(&model.Enemy{}).
		Find(&result)
	return result
}
