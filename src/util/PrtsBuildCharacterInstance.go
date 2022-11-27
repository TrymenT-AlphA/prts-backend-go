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

func PrtsBuildCharacterInstance(db *gorm.DB) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadFile(filepath.Join(cwd, "..", "src", "data", "CharacterInstance.json"))
	if err != nil {
		return err
	}
	var parser fastjson.Parser
	fjValue, err := parser.Parse(string(bytes))
	if err != nil {
		return err
	}

	fjValues := fjValue.GetArray()
	var characterInstances []model.CharacterInstance
	for _, fjValue := range fjValues {
		var characterInstance model.CharacterInstance
		characterInstance.CharacterID = string(fjValue.GetStringBytes("CharacterID"))
		characterInstance.Phase = fjValue.GetInt("Phase")
		characterInstance.Level = fjValue.GetInt("Level")
		characterInstance.RangeID = string(fjValue.GetStringBytes("RangeID"))
		characterInstance.MaxHp = fjValue.GetInt("MaxHp")
		characterInstance.Atk = fjValue.GetInt("Atk")
		characterInstance.Def = fjValue.GetInt("Def")
		characterInstance.MagicResistance = fjValue.GetFloat64("MagicResistance")
		characterInstance.Cost = fjValue.GetInt("Cost")
		characterInstance.BlockCnt = fjValue.GetInt("BlockCnt")
		characterInstance.MoveSpeed = fjValue.GetFloat64("MoveSpeed")
		characterInstance.AttackSpeed = fjValue.GetFloat64("AttackSpeed")
		characterInstance.BaseAttackTime = fjValue.GetFloat64("BaseAttackTime")
		characterInstance.RespawnTime = fjValue.GetInt("AttackSpeed")
		characterInstance.HpRecoveryPerSec = fjValue.GetFloat64("HpRecoveryPerSec")
		characterInstance.SpRecoveryPerSec = fjValue.GetFloat64("SpRecoveryPerSec")
		characterInstance.MaxDeployCount = fjValue.GetInt("MaxDeployCount")
		characterInstance.MaxDeckStackCnt = fjValue.GetInt("MaxDeckStackCnt")
		characterInstance.TauntLevel = fjValue.GetInt("TauntLevel")
		characterInstance.MassLevel = fjValue.GetInt("MassLevel")
		characterInstance.BaseForceLevel = fjValue.GetInt("BaseForceLevel")
		characterInstance.StunImmune = fjValue.GetBool("StunImmune")
		characterInstance.SilenceImmune = fjValue.GetBool("SilenceImmune")
		characterInstance.SleepImmune = fjValue.GetBool("SleepImmune")
		characterInstance.FrozenImmune = fjValue.GetBool("FrozenImmune")
		characterInstance.LevitateImmune = fjValue.GetBool("LevitateImmune")
		characterInstance.EvolveCost = string(fjValue.GetStringBytes("EvolveCost"))
		characterInstances = append(characterInstances, characterInstance)
	}

	err = db.
		Table("prts_characterinstance").
		Clauses(clause.OnConflict{UpdateAll: true}).
		CreateInBatches(&characterInstances, 500).
		Error
	if err != nil {
		return err
	}

	return nil
}
