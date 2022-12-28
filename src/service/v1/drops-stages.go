package apiv1

import (
	"prts-backend/src/model"
	"prts-backend/src/service"
)

func DropsStages() ([]model.Stage, error) {
	var result []model.Stage
	err := service.DB.
		Raw(`SELECT prts_Stage.*
			FROM prts_Stage 
			JOIN prts_Drop
			ON prts_Stage.Difficulty != ?
			AND prts_Stage.Id = prts_Drop.StageId
			GROUP BY Id`, "FOUR_STAR").
		Scan(&result).
		Error
	return result, err
}
