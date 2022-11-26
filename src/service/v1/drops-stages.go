package service

import "prts-backend/src/model"

func DropsStages() []model.Stage {
	var result []model.Stage
	db.
		Model(&model.Stage{}).
		Having("COUNT(Drops) >= 1").
		Find(&result)
	return result
}
