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

func PrtsBuildEI_S(db *gorm.DB) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadFile(filepath.Join(cwd, "..", "src", "data", "EI_S.json"))
	if err != nil {
		return err
	}
	var parser fastjson.Parser
	fjValue, err := parser.Parse(string(bytes))
	if err != nil {
		return err
	}

	fjValues := fjValue.GetArray()
	var ei_ss []model.EI_S
	for _, fjValue := range fjValues {
		var ei_s model.EI_S
		ei_s.EnemyInstanceEnemyID = string(fjValue.GetStringBytes("EnemyID"))
		ei_s.EnemyInstanceLevel = fjValue.GetInt("EnemyInstanceLevel")
		ei_s.StageID = string(fjValue.GetStringBytes("StageID"))
		ei_ss = append(ei_ss, ei_s)
	}
	//! for test use, ignore some bad data
	// for _, ei_s := range ei_ss {
	// 	if err = db.Table("ei_s").Clauses(clause.OnConflict{UpdateAll: true}).Create(&ei_s).Error; err != nil {
	// 		log.Print(err) // ignore some bad data
	// 	}
	// }

	err = db.
		Table("prts_ei_s").
		Clauses(clause.OnConflict{UpdateAll: true}).
		Create(&ei_ss).
		Error
	if err != nil {
		return err
	}

	return nil
}
