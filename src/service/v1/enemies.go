package service

import "prts-backend/src/model"

func Enemies() []model.Enemy {
	var result []model.Enemy
	db.
		Model(&model.Enemy{}).
		Find(&result)
	return result
}
