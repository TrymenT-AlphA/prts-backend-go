package service

import "prts-backend/src/model"

func DropsStages() []model.Stage {
	var result []model.Stage
	db.
		Raw(`SELECT prts_stage.*
			FROM prts_stage
			JOIN prts_drop
			ON prts_stage.ID = prts_drop.StageID
			GROUP BY ID`).
		Scan(&result)
	return result
}
