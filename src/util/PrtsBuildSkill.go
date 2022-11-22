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

func PrtsBuildSkill(db *gorm.DB) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadFile(filepath.Join(cwd, "..", "src", "data", "Skill.json"))
	if err != nil {
		return err
	}
	var parser fastjson.Parser
	fjValue, err := parser.Parse(string(bytes))
	if err != nil {
		return err
	}

	fjValues := fjValue.GetArray()
	var skills []model.Skill
	for _, fjValue := range fjValues {
		var skill model.Skill
		skill.ID = string(fjValue.GetStringBytes("ID"))
		skill.IconID = string(fjValue.GetStringBytes("IconID"))
		skills = append(skills, skill)
	}

	err = db.
		Table("skill").
		Clauses(clause.OnConflict{UpdateAll: true}).
		Create(&skills).
		Error
	if err != nil {
		return err
	}

	return nil
}
