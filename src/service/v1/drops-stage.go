package service

import "prts-backend/src/model"

func DropsStage(id string) model.Stage {
	var result model.Stage
	db.
		Model(&model.Stage{}).
		Where(&model.Stage{ID: id}).
		Limit(1).
		Preload("Drops").
		Preload("Drops.Item").
		Preload("Drops.Stage").
		Find(&result)
	return result
}
