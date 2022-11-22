package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"prts-backend/src/model"

	"github.com/valyala/fastjson"
	"gorm.io/gorm"
)

func PrtsBuildSkillInstance(db *gorm.DB) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadFile(filepath.Join(cwd, "..", "src", "data", "SkillInstance.json"))
	if err != nil {
		return err
	}
	var parser fastjson.Parser
	fjValue, err := parser.Parse(string(bytes))
	if err != nil {
		return err
	}

	fjValues := fjValue.GetArray()
	var skillInstances []model.SkillInstance
	for _, fjValue := range fjValues {
		var skillInstance model.SkillInstance
		skillInstance.SkillID = string(fjValue.GetStringBytes("SkillID"))
		skillInstance.Level = fjValue.GetInt("Level")
		skillInstance.Name = string(fjValue.GetStringBytes("Name"))
		skillInstance.RangeID = string(fjValue.GetStringBytes("RangeID"))
		skillInstance.Description = string(fjValue.GetStringBytes("Description"))
		skillInstance.Type = fjValue.GetInt("Type")
		skillInstance.DurationType = fjValue.GetInt("DurationType")
		skillInstance.SpType = fjValue.GetInt("SpType")
		skillInstance.MaxChargeTime = fjValue.GetInt("MaxChargeTime")
		skillInstance.SpCost = fjValue.GetInt("SpCost")
		skillInstance.InitSp = fjValue.GetInt("InitSp")
		skillInstance.Duration = fjValue.GetInt("Duration")
		skillInstances = append(skillInstances, skillInstance)
	}
	db.Create(&skillInstances)
	return nil
}
