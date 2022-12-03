package apiv1

import (
	"prts-backend/src/model"
	"prts-backend/src/service"
)

func DropsStage(id string) model.Stage {
	var result model.Stage
	service.DB.
		Model(&model.Stage{}).
		Where(&model.Stage{Id: id}).
		Limit(1).
		Preload("Drops").
		Preload("Drops.Item").
		Preload("Drops.Stage").
		Find(&result)
	return result
}
