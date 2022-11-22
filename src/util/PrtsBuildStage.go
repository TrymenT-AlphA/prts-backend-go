package util

import (
	"database/sql"
	"io/ioutil"
	"os"
	"path/filepath"
	"prts-backend/src/model"

	"github.com/ahmetb/go-linq/v3"
	"github.com/valyala/fastjson"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func PrtsBuildStage(db *gorm.DB) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadFile(filepath.Join(cwd, "..", "src", "data", "Stage.json"))
	if err != nil {
		return err
	}
	var parser fastjson.Parser
	fjValue, err := parser.Parse(string(bytes))
	if err != nil {
		return err
	}

	fjValues := fjValue.GetArray()
	var stages []model.Stage
	for _, fjValue := range fjValues {
		var stage model.Stage
		stage.ID = string(fjValue.GetStringBytes("ID"))
		stage.Type = string(fjValue.GetStringBytes("Type"))
		stage.Difficulty = string(fjValue.GetStringBytes("Difficulty"))
		stage.PerformanceStageFlag = string(fjValue.GetStringBytes("PerformanceStageFlag"))
		stage.DiffGroup = string(fjValue.GetStringBytes("DiffGroup"))
		stage.UnlockCondition = string(fjValue.GetStringBytes("UnlockCondition"))
		stage.LevelID = string(fjValue.GetStringBytes("LevelID"))
		stage.ZoneID = string(fjValue.GetStringBytes("ZoneID"))
		stage.Code = string(fjValue.GetStringBytes("Code"))
		stage.Name = string(fjValue.GetStringBytes("Name"))
		stage.Description = string(fjValue.GetStringBytes("Description"))
		stage.DangerLevel = string(fjValue.GetStringBytes("DangerLevel"))
		stage.CanPractice = fjValue.GetBool("CanPractice")
		stage.PracticeTicketCost = fjValue.GetInt("PracticeTicketCost")
		stage.ApCost = fjValue.GetInt("ApCost")
		HardStageID := string(fjValue.GetStringBytes("HardStageID"))
		if HardStageID == "" {
			stage.HardStageID = sql.NullString{String: HardStageID, Valid: false}
		} else {
			stage.HardStageID = sql.NullString{String: HardStageID, Valid: true}
		}
		stages = append(stages, stage)
	}
	// var sortedStages []model.Stage
	linq.From(stages).Sort(func(i interface{}, j interface{}) bool {
		if !i.(model.Stage).HardStageID.Valid {
			return true
		} else if !j.(model.Stage).HardStageID.Valid {
			return false
		} else {
			return true
		}
	}).ToSlice(&stages)

	err = db.
		Table("stage").
		Clauses(clause.OnConflict{UpdateAll: true}).
		Create(&stages).
		Error
	if err != nil {
		return err
	}

	return nil
}
