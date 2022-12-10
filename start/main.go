package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"prts-backend/src/server"

	"github.com/gofiber/fiber/v2"
	"github.com/tidwall/gjson"
)

func main() {
	// simply set prod option
	prod := flag.Bool("prod", false, "prod")
	flag.Parse()
	var configFile string
	if *prod { // if you use prod option
		configFile = "config.prod.json"
	} else { // default
		configFile = "config.dev.json"
	}
	// get cwd
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("[INFO] Reading Config from " + configFile)
	}
	// read config file
	bytes, err := os.ReadFile(filepath.Join(cwd, "..", "start", configFile))
	if err != nil {
		log.Fatal(err)
	}
	config := gjson.ParseBytes(bytes)
	// run the server
	if err = server.Run(
		config.Get("port").String(),
		config.Get("dsn").String(),
		fiber.Config{
			Prefork:       config.Get("prefork").Bool(),
			CaseSensitive: config.Get("caseSensitive").Bool(),
			StrictRouting: config.Get("strictRouting").Bool(),
			ServerHeader:  config.Get("serverHeader").String(),
			AppName:       config.Get("appName").String(),
		}); err != nil {
		log.Fatal(err)
	}
}
