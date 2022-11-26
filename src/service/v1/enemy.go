package service

import "prts-backend/src/model"

func Enemy(id string) model.Enemy {
	var result model.Enemy
	db.
		Model(&model.Enemy{}).
		Where(&model.Enemy{ID: id}).
		Limit(1).
		Find(&result)
	return result
}
