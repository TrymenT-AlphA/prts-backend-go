package apiv1

import (
	"prts-backend/src/model"
	"prts-backend/src/service"
)

func Stages() []model.Stage {
	var result []model.Stage
	service.DB.
		Model(&model.Stage{}).
		Where("Difficulty != ?", "FOUR_STAR").
		Find(&result)
	return result
}

func MainStages() []model.Stage {
	var result []model.Stage
	service.DB.
		Model(&model.Stage{}).
		Where("Difficulty != ?", "FOUR_STAR").
		Where(&model.Stage{
			Type: "GUID",
		}).
		Or(&model.Stage{
			Type: "SUB",
		}).
		Or(&model.Stage{
			Type: "SPECIAL_STORY",
		}).
		Or(&model.Stage{
			Type: "MAIN",
		}).
		Find(&result)
	return result
}

func PermStages() []model.Stage {
	var result []model.Stage
	service.DB.
		Model(&model.Stage{}).
		Where("Difficulty != ?", "FOUR_STAR").
		Where(&model.Stage{
			Type: "CAMPAIGN",
		}).
		Or(&model.Stage{
			Type: "CLIMB_TOWER",
		}).
		Or(&model.Stage{
			Type: "DAILY",
		}).
		Find(&result)
	return result
}

func ActiStages() []model.Stage {
	var result []model.Stage
	service.DB.
		Model(&model.Stage{}).
		Where("Difficulty != ?", "FOUR_STAR").
		Where(&model.Stage{
			Type: "ACTIVITY",
		}).
		Find(&result)
	return result
}
