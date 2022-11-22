package service

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"prts-backend/src/util"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB = nil

func InitDatabase() error {
	if db != nil {
		return nil
	}
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	dsn, err := ioutil.ReadFile(filepath.Join(cwd, "..", ".env"))
	if err != nil {
		return err
	}

	db, err = gorm.Open(mysql.Open(string(dsn)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return err
	}

	err = util.PrtsAutoMigrate(db)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("[SUCCESS] auto migrate database")
	}
	err = util.PrtsAutoBuild(db)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("[SUCCESS] auto build table data")
	}

	return nil
}
