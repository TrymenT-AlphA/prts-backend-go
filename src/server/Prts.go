package server

import (
	"log"
	"prts-backend/src/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Prts(port string) {

	err := service.InitDatabase(false)
	if err != nil {
		log.Fatal(err)
	}

	prts := fiber.New()

	prts.Use(cors.New())

	prts.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("prts-backend is running!")
	})

	prts.Get("/drops/stages/:StageID", func(c *fiber.Ctx) error {
		return c.JSON(service.DropsStage(c.Params("StageID")))
	})

	prts.Listen(port)
}
