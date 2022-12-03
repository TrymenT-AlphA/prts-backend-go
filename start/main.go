package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"prts-backend/src/server"

	"github.com/gofiber/fiber/v2"
	"github.com/tidwall/gjson"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := ioutil.ReadFile(filepath.Join(cwd, "..", "start", "config.json"))
	if err != nil {
		log.Fatal(err)
	}
	config := gjson.ParseBytes(bytes)
	server.Run(config.Get("port").String(), config.Get("dsn").String(), fiber.Config{
		Prefork:       config.Get("prefork").Bool(),
		CaseSensitive: config.Get("caseSensitive").Bool(),
		StrictRouting: config.Get("strictRouting").Bool(),
		ServerHeader:  config.Get("serverHeader").String(),
		AppName:       config.Get("appName").String(),
	})
}
