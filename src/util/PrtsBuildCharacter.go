package util

import (
	"database/sql"
	"io/ioutil"
	"os"
	"path/filepath"
	"prts-backend/src/model"

	"github.com/ahmetb/go-linq/v3"
	"github.com/valyala/fastjson"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func PrtsBuildCharacter(db *gorm.DB) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadFile(filepath.Join(cwd, "..", "src", "data", "Character.json"))
	if err != nil {
		return err
	}
	var parser fastjson.Parser
	fjValue, err := parser.Parse(string(bytes))
	if err != nil {
		return err
	}

	fjValues := fjValue.GetArray()
	var characters []model.Character
	for _, fjValue := range fjValues {
		var character model.Character
		character.ID = string(fjValue.GetStringBytes("ID"))
		character.Name = string(fjValue.GetStringBytes("Name"))
		character.Description = string(fjValue.GetStringBytes("Description"))
		character.CanUseGeneralPotentialItem = fjValue.GetBool("CanUseGeneralPotentialItem")
		character.CanUseActivityPotentialItem = fjValue.GetBool("CanUseActivityPotentialItem")
		character.PotentialItemID = string(fjValue.GetStringBytes("PotentialItemID"))
		character.ActivityPotentialItemID = string(fjValue.GetStringBytes("ActivityPotentialItemID"))
		character.NationID = string(fjValue.GetStringBytes("NationID"))
		character.GroupID = string(fjValue.GetStringBytes("GroupID"))
		character.TeamID = string(fjValue.GetStringBytes("TeamID"))
		character.DisplayNumber = string(fjValue.GetStringBytes("DisplayNumber"))
		character.Appellation = string(fjValue.GetStringBytes("Appellation"))
		character.Position = string(fjValue.GetStringBytes("Position"))
		character.TagList = string(fjValue.GetStringBytes("TagList"))
		character.ItemUsage = string(fjValue.GetStringBytes("ItemUsage"))
		character.ItemDesc = string(fjValue.GetStringBytes("ItemDesc"))
		character.ItemObtainApproach = string(fjValue.GetStringBytes("ItemDesc"))
		character.IsNotObtainable = fjValue.GetBool("IsNotObtainable")
		character.IsSpChar = fjValue.GetBool("IsSpChar")
		character.MaxPotentialLevel = fjValue.GetInt("MaxPotentialLevel")
		character.Rarity = fjValue.GetInt("Rarity")
		character.Profession = string(fjValue.GetStringBytes("Profession"))
		character.SubProfessionID = string(fjValue.GetStringBytes("SubProfessionID"))
		character.AllSkillLvlupList = string(fjValue.GetStringBytes("AllSkillLvlupList"))
		character.PotentialList = string(fjValue.GetStringBytes("PotentialList"))
		tokenIDString := string(fjValue.GetStringBytes("TokenID"))
		if tokenIDString == "" {
			character.TokenID = sql.NullString{String: tokenIDString, Valid: false}
		} else {
			character.TokenID = sql.NullString{String: tokenIDString, Valid: true}
		}
		characters = append(characters, character)
	}
	linq.From(characters).Sort(func(i interface{}, j interface{}) bool {
		if !i.(model.Character).TokenID.Valid {
			return true
		} else if !j.(model.Character).TokenID.Valid {
			return false
		} else {
			return true
		}
	}).ToSlice(&characters)

	err = db.
		Table("prts_character").
		Clauses(clause.OnConflict{UpdateAll: true}).
		Create(&characters).
		Error
	if err != nil {
		return err
	}

	return nil
}
