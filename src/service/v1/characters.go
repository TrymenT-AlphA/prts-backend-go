package service

import "prts-backend/src/model"

func Characters() []model.Character {
	var result []model.Character
	db.
		Model(&model.Character{}).
		Find(&result)
	return result
}
