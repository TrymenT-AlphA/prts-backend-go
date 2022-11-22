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

func PrtsBuildBuildingSkill(db *gorm.DB) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadFile(filepath.Join(cwd, "..", "src", "data", "BuildingSkill.json"))
	if err != nil {
		return err
	}
	var parser fastjson.Parser
	fjValue, err := parser.Parse(string(bytes))
	if err != nil {
		return err
	}

	fjValues := fjValue.GetArray()
	var buildingSkills []model.BuildingSkill
	for _, fjValue := range fjValues {
		var buildingSkill model.BuildingSkill
		buildingSkill.ID = string(fjValue.GetStringBytes("ID"))
		buildingSkill.Name = string(fjValue.GetStringBytes("Name"))
		buildingSkill.Icon = string(fjValue.GetStringBytes("Icon"))
		buildingSkill.SortID = fjValue.GetInt("SortID")
		buildingSkill.Category = string(fjValue.GetStringBytes("Category"))
		buildingSkill.RoomType = string(fjValue.GetStringBytes("RoomType"))
		buildingSkill.Description = string(fjValue.GetStringBytes("Description"))
		buildingSkills = append(buildingSkills, buildingSkill)
	}
	if err = db.Table("buildingskill").Clauses(clause.OnConflict{UpdateAll: true}).Create(&buildingSkills).Error; err != nil {
		return err
	}
	return nil
}
