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

func BuildTalent(db *gorm.DB) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadFile(filepath.Join(cwd, "..", "src", "data", "Talent.json"))
	if err != nil {
		return err
	}
	var parser fastjson.Parser
	fjValue, err := parser.Parse(string(bytes))
	if err != nil {
		return err
	}

	fjValues := fjValue.GetArray()
	var talents []model.Talent
	for _, fjValue := range fjValues {
		var talent model.Talent
		talent.PrefabKey = fjValue.GetInt("PrefabKey")
		talent.Phase = fjValue.GetInt("Phase")
		talent.Level = fjValue.GetInt("Level")
		talent.PotentialRank = fjValue.GetInt("PotentialRank")
		talent.Name = string(fjValue.GetStringBytes("Name"))
		talent.Description = string(fjValue.GetStringBytes("Description"))
		talent.RangeID = string(fjValue.GetStringBytes("RangeID"))
		talent.CharacterID = string(fjValue.GetStringBytes("CharacterID"))
		talents = append(talents, talent)
	}

	err = db.
		Model(&model.Talent{}).
		Clauses(clause.OnConflict{UpdateAll: true}).
		Create(&talents).
		Error
	if err != nil {
		return err
	}

	return nil
}
