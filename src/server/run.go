package server

import (
	"prts-backend/src/service"
	apiv1 "prts-backend/src/service/v1"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Run(port string, dsn string, config fiber.Config) error {
	if !fiber.IsChild() {
		if err := service.InitDatabase(dsn); err != nil {
			return err
		}
	}

	// the very basic server function
	app := fiber.New(config)
	app.Use(cors.New())
	app.Use(recover.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("prts-backend service is running!")
	})
	// mount api
	app.Mount("/api/v1", apiv1.GetAPI())
	// listen to port
	err := app.Listen(port)
	if err != nil {
		return err
	}

	return nil
}
