package service

import "prts-backend/src/model"

func Character(id string) model.Character {
	var result model.Character
	db.
		Model(&model.Character{}).
		Where(&model.Character{ID: id}).
		Limit(1).
		Find(&result)
	return result
}
