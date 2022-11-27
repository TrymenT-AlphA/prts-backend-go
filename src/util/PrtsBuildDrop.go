package util

import (
	"database/sql"
	"io/ioutil"
	"os"
	"path/filepath"
	"prts-backend/src/model"
	"time"

	"github.com/valyala/fastjson"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func PrtsBuildDrop(db *gorm.DB) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadFile(filepath.Join(cwd, "..", "src", "data", "Drop.json"))
	if err != nil {
		return err
	}
	var parser fastjson.Parser
	fjValue, err := parser.Parse(string(bytes))
	if err != nil {
		return err
	}

	fjValues := fjValue.GetArray()
	var drops []model.Drop
	timeLayout := "2006-01-02T15:04:05Z"
	for _, fjValue := range fjValues {
		var drop model.Drop
		drop.ItemID = string(fjValue.GetStringBytes("ItemID"))
		drop.StageID = string(fjValue.GetStringBytes("StageID"))
		drop.Times = fjValue.GetInt("Times")
		drop.Quantity = fjValue.GetInt("Quantity")
		drop.StdDev = fjValue.GetFloat64("StdDev")
		drop.Start, err = time.Parse(timeLayout, string(fjValue.GetStringBytes("Start")))
		if err != nil {
			return err
		}
		endString := string(fjValue.GetStringBytes("End"))
		if endString == "" {
			drop.End = sql.NullTime{Time: time.Now(), Valid: false}
		} else {
			endTime, err := time.Parse(timeLayout, string(fjValue.GetStringBytes("End")))
			if err != nil {
				return err
			}
			drop.End = sql.NullTime{Time: endTime, Valid: true}
		}
		if err != nil {
			return err
		}
		drop.UpdateAt = time.Now()
		drops = append(drops, drop)
	}
	//! for test use, ignore some bad data
	// for _, drop := range drops {
	// 	if err = db.Table("drop").Clauses(clause.OnConflict{UpdateAll: true}).Create(&drop).Error; err != nil {
	// 		log.Print(err) // ignore some bad data
	// 	}
	// }

	err = db.
		Table("prts_drop").
		Clauses(clause.OnConflict{UpdateAll: true}).
		Create(&drops).
		Error
	if err != nil {
		return err
	}

	return nil
}
