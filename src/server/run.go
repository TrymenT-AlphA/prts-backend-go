package server

import (
	"log"
	"prts-backend/src/service"
	apiv1 "prts-backend/src/service/v1"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Run(port string, dsn string, config fiber.Config) {
	if !fiber.IsChild() {
		err := service.InitDatabase(dsn, false)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := service.InitDatabase(dsn, true)
		if err != nil {
			log.Fatal(err)
		}
	}

	app := fiber.New(config)

	app.Use(cors.New())

	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("prts-backend service is running!")
	})

	app.Mount("/api/v1", apiv1.GetAPI())

	err := app.Listen(port)
	if err != nil {
		log.Fatal(err)
	}
}
