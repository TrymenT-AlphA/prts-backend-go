package apiv1

import (
	"prts-backend/src/model"
	"prts-backend/src/service"
)

func Character(id string) model.Character {
	var result model.Character
	service.DB.
		Model(&model.Character{}).
		Where(&model.Character{Id: id}).
		Limit(1).
		Find(&result)
	return result
}
