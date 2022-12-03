package apiv1

import (
	"prts-backend/src/model"
	"prts-backend/src/service"
)

func Stages() []model.Stage {
	var result []model.Stage
	service.DB.
		Model(&model.Stage{}).
		Find(&result)
	return result
}
