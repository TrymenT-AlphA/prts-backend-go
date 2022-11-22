package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"prts-backend/src/model"

	"github.com/valyala/fastjson"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func PrtsBuildC_BS(db *gorm.DB) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadFile(filepath.Join(cwd, "..", "src", "data", "C_BS.json"))
	if err != nil {
		return err
	}
	var parser fastjson.Parser
	fjValue, err := parser.Parse(string(bytes))
	if err != nil {
		return err
	}

	fjValues := fjValue.GetArray()
	var c_bss []model.C_BS
	for _, fjValue := range fjValues {
		var c_bs model.C_BS
		c_bs.CharacterID = string(fjValue.GetStringBytes("CharacterID"))
		c_bs.BuildingSkillID = string(fjValue.GetStringBytes("BuildingSkillID"))
		c_bs.Phase = fjValue.GetInt("Phase")
		c_bs.Level = fjValue.GetInt("Level")
		c_bss = append(c_bss, c_bs)
	}

	err = db.
		Table("c_bs").
		Clauses(clause.OnConflict{UpdateAll: true}).
		Create(&c_bss).
		Error
	if err != nil {
		return err
	}

	return nil
}
