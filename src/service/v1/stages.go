package service

import "prts-backend/src/model"

func Stages() []model.Stage {
	var result []model.Stage
	db.
		Model(&model.Stage{}).
		Find(&result)
	return result
}
