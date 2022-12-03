package apiv1

import (
	"prts-backend/src/model"
	"prts-backend/src/service"
)

func DropsStages() ([]model.Stage, error) {
	var result []model.Stage
	err := service.DB.
		Raw(`SELECT prts_stage.*
			FROM prts_stage
			JOIN prts_drop
			ON prts_stage.ID = prts_drop.StageID
			GROUP BY ID`).
		Scan(&result).
		Error
	return result, err
}
