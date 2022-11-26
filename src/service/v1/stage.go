package service

import "prts-backend/src/model"

func Stage(id string) model.Stage {
	var result model.Stage
	db.
		Model(&model.Stage{}).
		Where(&model.Stage{ID: id}).
		Limit(1).
		Find(&result)
	return result
}
