package server

import (
	"log"

	"github.com/gofiber/fiber"
)

func Prts(port string) {
	err := InitDB()
	if err != nil {
		log.Fatal(err)
	}

	prts := fiber.New()

	prts.Get("/", func(c *fiber.Ctx) {
		c.SendString("prts-backend running!")
	})

	prts.Listen(port)
}
