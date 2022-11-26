package service

import "prts-backend/src/model"

func Drops() []model.Drop {
	var result []model.Drop
	db.
		Model(&model.Drop{}).
		Find(&result)
	return result
}
