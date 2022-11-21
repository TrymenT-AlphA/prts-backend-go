package test

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var TestDSN = "chongkai:123456@tcp(127.0.0.1:3306)/prts-go-test?charset=utf8mb4&parseTime=True&loc=Local"
var TestDB *gorm.DB = nil

func InitTestDB() error {
	if TestDB != nil {
		return nil
	}
	var err error
	TestDB, err = gorm.Open(mysql.Open(TestDSN), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
