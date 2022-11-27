package test

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var dsn = "chongkai:123456@tcp(127.0.0.1:3306)/prts-go-test?charset=utf8mb4&parseTime=True&loc=Local"
var db *gorm.DB = nil

func InitTestDB() error {
	if db != nil {
		return nil
	}
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "prts_",
			SingularTable: true,
			NoLowerCase:   true,
		},
		//* enable foreignKey to pass test
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	if err != nil {
		return err
	}
	return nil
}
