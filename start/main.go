package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"prts-backend/src/server"

	"github.com/gofiber/fiber/v2"
	"github.com/tidwall/gjson"
)

var prod = flag.Bool("prod", false, "prod")

func main() {
	flag.Parse()
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	var bytes []byte
	if *prod {
		bytes, err = ioutil.ReadFile(filepath.Join(cwd, "..", "start", "config.prod.json"))
		log.Print("[INFO] Reading Config from config.prod.json")
	} else {
		bytes, err = ioutil.ReadFile(filepath.Join(cwd, "..", "start", "config.dev.json"))
		log.Print("[INFO] Reading Config from config.dev.json")
	}

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
