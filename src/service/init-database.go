package service

import (
	"log"
	"prts-backend/src/util"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB = nil

func InitDatabase(dsn string, isChild bool) error {
	if db != nil {
		return nil
	}

	db, err := gorm.Open(mysql.Open(string(dsn)), &gorm.Config{
		CreateBatchSize: 500,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "prts_",
			SingularTable: true,
			NoLowerCase:   true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return err
	}

	if !isChild {
		err = util.AutoMigrate(db)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Print("[SUCCESS] auto migrate database")
		}
		err = util.AutoBuild(db)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Print("[SUCCESS] auto build table data")
		}
	}

	return nil
}
