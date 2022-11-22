package server

import (
	"log"
	"prts-backend/src/service"

	"github.com/gofiber/fiber"
)

func Prts(port string) {

	err := service.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}

	prts := fiber.New()

	prts.Get("/", func(c *fiber.Ctx) {
		c.SendString("prts-backend is running!")
	})

	prts.Listen(port)
}
