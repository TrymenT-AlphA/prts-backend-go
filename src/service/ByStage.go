package service

import "prts-backend/src/model"

func ByStage(StageID string) model.Stage {
	var stage model.Stage
	db.
		Model(&model.Stage{}).
		Where(&model.Stage{ID: StageID}).
		Preload("Drops").
		Preload("Drops.Item").
		Preload("Drops.Stage").
		First(&stage)
	return stage
}
