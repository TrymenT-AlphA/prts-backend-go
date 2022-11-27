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

func BuildEnemy(db *gorm.DB) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadFile(filepath.Join(cwd, "..", "src", "data", "Enemy.json"))
	if err != nil {
		return err
	}
	var parser fastjson.Parser
	fjValue, err := parser.Parse(string(bytes))
	if err != nil {
		return err
	}

	fjValues := fjValue.GetArray()
	var enemies []model.Enemy
	for _, fjValue := range fjValues {
		var enemy model.Enemy
		enemy.ID = string(fjValue.GetStringBytes("ID"))
		enemy.SortID = fjValue.GetInt("SortID")
		enemy.Name = string(fjValue.GetStringBytes("Name"))
		enemy.Index = string(fjValue.GetStringBytes("Index"))
		enemy.Race = string(fjValue.GetStringBytes("Race"))
		enemy.Level = string(fjValue.GetStringBytes("Level"))
		enemy.AttackType = string(fjValue.GetStringBytes("AttackType"))
		enemy.Endure = string(fjValue.GetStringBytes("Endure"))
		enemy.Attack = string(fjValue.GetStringBytes("Attack"))
		enemy.Defence = string(fjValue.GetStringBytes("Defence"))
		enemy.Resistance = string(fjValue.GetStringBytes("Resistance"))
		enemy.Description = string(fjValue.GetStringBytes("Description"))
		enemy.Ability = string(fjValue.GetStringBytes("Ability"))
		enemy.IsInvalidKilled = fjValue.GetBool("IsInvalidKilled")
		enemies = append(enemies, enemy)
	}

	err = db.
		Model(&model.Enemy{}).
		Clauses(clause.OnConflict{UpdateAll: true}).
		Create(&enemies).
		Error
	if err != nil {
		return err
	}

	return nil
}
