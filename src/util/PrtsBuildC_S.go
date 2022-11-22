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

func PrtsBuildC_S(db *gorm.DB) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadFile(filepath.Join(cwd, "..", "src", "data", "C_S.json"))
	if err != nil {
		return err
	}
	var parser fastjson.Parser
	fjValue, err := parser.Parse(string(bytes))
	if err != nil {
		return err
	}

	fjValues := fjValue.GetArray()
	var c_ss []model.C_S
	for _, fjValue := range fjValues {
		var c_s model.C_S
		c_s.CharacterID = string(fjValue.GetStringBytes("CharacterID"))
		c_s.SkillID = string(fjValue.GetStringBytes("SkillID"))
		c_s.LvlupCostCond = string(fjValue.GetStringBytes("LvlupCostCond"))
		c_ss = append(c_ss, c_s)
	}
	if err = db.Table("c_s").Clauses(clause.OnConflict{UpdateAll: true}).Create(&c_ss).Error; err != nil {
		return err
	}
	return nil
}
