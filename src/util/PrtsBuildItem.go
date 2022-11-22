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

func PrtsBuildItem(db *gorm.DB) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadFile(filepath.Join(cwd, "..", "src", "data", "Item.json"))
	if err != nil {
		return err
	}
	var parser fastjson.Parser
	fjValue, err := parser.Parse(string(bytes))
	if err != nil {
		return err
	}

	fjValues := fjValue.GetArray()
	var items []model.Item
	for _, fjValue := range fjValues {
		var item model.Item
		item.ID = string(fjValue.GetStringBytes("ID"))
		item.Name = string(fjValue.GetStringBytes("Name"))
		item.Description = string(fjValue.GetStringBytes("Description"))
		item.Rarity = fjValue.GetInt("Rarity")
		item.IconID = string(fjValue.GetStringBytes("IconID"))
		item.SortID = fjValue.GetInt("SortID")
		item.Usage = string(fjValue.GetStringBytes("Usage"))
		item.ObtainApproach = string(fjValue.GetStringBytes("ObtainApproach"))
		item.ClassifyType = string(fjValue.GetStringBytes("ClassifyType"))
		item.Type = string(fjValue.GetStringBytes("Type"))
		items = append(items, item)
	}

	err = db.
		Table("item").
		Clauses(clause.OnConflict{UpdateAll: true}).
		Create(&items).
		Error
	if err != nil {
		return err
	}

	return nil
}
