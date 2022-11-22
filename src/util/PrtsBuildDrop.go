package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"prts-backend/src/model"
	"time"

	"github.com/valyala/fastjson"
	"gorm.io/gorm"
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
			return nil
		}
		drop.End, err = time.Parse(timeLayout, string(fjValue.GetStringBytes("End")))
		if err != nil {
			return nil
		}
		drop.UpdateAt = time.Now()
		drops = append(drops, drop)
	}
	db.Create(&drops)
	return nil
}
