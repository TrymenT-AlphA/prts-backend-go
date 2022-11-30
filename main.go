package main

import (
	"prts-backend/src/server"

	"github.com/gofiber/fiber/v2"
)

func main() {
	var port = ":33333"
	server.Run(&port, &fiber.Config{
		Prefork:       false,
		CaseSensitive: false,
		StrictRouting: false,
		ServerHeader:  "prts-backend powered by fiber",
		AppName:       "prts-backend v1",
	})
}
