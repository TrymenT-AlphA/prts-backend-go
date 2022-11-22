package main

// prts-backend
// using fiber gorm fastjson for better performace

import (
	"log"

	"github.com/valyala/fastjson"
)

// "log"

// "gorm.io/driver/mysql"
// "gorm.io/gorm"

func main() {
	fjValue, _ := fastjson.Parse(`{
		"s": null,
	}`)
	s := string(fjValue.GetStringBytes("s"))
	log.Print(s)
}
