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

func PrtsBuildEnemyInstance(db *gorm.DB) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadFile(filepath.Join(cwd, "..", "src", "data", "EnemyInstance.json"))
	if err != nil {
		return err
	}
	var parser fastjson.Parser
	fjValue, err := parser.Parse(string(bytes))
	if err != nil {
		return err
	}

	fjValues := fjValue.GetArray()
	var enemyInstances []model.EnemyInstance
	for _, fjValue := range fjValues {
		var enemyInstance model.EnemyInstance
		enemyInstance.EnemyID = string(fjValue.GetStringBytes("EnemyID"))
		enemyInstance.Level = fjValue.GetInt("Level")
		enemyInstance.MaxHp = fjValue.GetInt("MaxHp")
		enemyInstance.Atk = fjValue.GetInt("Atk")
		enemyInstance.Def = fjValue.GetInt("Def")
		enemyInstance.MagicResistance = fjValue.GetFloat64("MagicResistance")
		enemyInstance.LifePointReduce = fjValue.GetInt("LifePointReduce")
		enemyInstance.AttackSpeed = fjValue.GetFloat64("AttackSpeed")
		enemyInstance.BaseAttackTime = fjValue.GetFloat64("BaseAttackTime")
		enemyInstance.RangeRadius = fjValue.GetFloat64("RangeRadius")
		enemyInstance.MoveSpeed = fjValue.GetFloat64("MoveSpeed")
		enemyInstance.MassLevel = fjValue.GetInt("MassLevel")
		enemyInstance.RespawnTime = fjValue.GetFloat64("RespawnTime")
		enemyInstance.SilenceImmune = fjValue.GetBool("SilenceImmune")
		enemyInstance.StunImmune = fjValue.GetBool("StunImmune")
		enemyInstance.SleepImmune = fjValue.GetBool("SleepImmune")
		enemyInstance.FrozenImmune = fjValue.GetBool("FrozenImmune")
		enemyInstance.LevitateImmune = fjValue.GetBool("LevitateImmune")
		enemyInstances = append(enemyInstances, enemyInstance)
	}
	if err = db.Table("enemyinstance").Clauses(clause.OnConflict{UpdateAll: true}).Create(&enemyInstances).Error; err != nil {
		return err
	}
	return nil
}
