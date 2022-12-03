package apiv1

import (
	"prts-backend/src/model"
	"prts-backend/src/service"
)

func Characters() []model.Character {
	var result []model.Character
	service.DB.
		Model(&model.Character{}).
		Find(&result)
	return result
}
