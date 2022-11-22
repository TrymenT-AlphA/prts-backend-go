package server

import (
	"log"
	"prts-backend/src/service"

	"github.com/gofiber/fiber/v2"
)

func Prts(port string) {

	err := service.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}

	prts := fiber.New()

	prts.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("prts-backend is running!")
	})

	prts.Get("/byStage/:StageID", func(c *fiber.Ctx) error {
		return c.JSON(service.ByStage(c.Params("StageID")))
	})

	prts.Listen(port)
}
