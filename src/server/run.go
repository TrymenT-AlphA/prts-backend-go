package server

import (
	"log"
	"prts-backend/src/service/v1"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Run(port *string, config *fiber.Config) {

	err := service.InitDatabase(false)
	if err != nil {
		log.Fatal(err)
	}

	var app *fiber.App
	if config != nil {
		app = fiber.New(*config)
	} else {
		app = fiber.New(fiber.Config{
			Prefork:       true,
			CaseSensitive: false,
			StrictRouting: false,
			ServerHeader:  "prts-backend powered by fiber",
			AppName:       "prts-backend v1",
		})
	}

	app.Use(cors.New())

	app.Use(recover.New())

	api := app.Group("/api")

	v1 := api.Group("/v1")

	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("prts-backend v1 is running!")
	})

	v1.Get("/characters", func(c *fiber.Ctx) error {
		return c.JSON(service.Characters())
	})

	v1.Get("/characters/:id", func(c *fiber.Ctx) error {
		return c.JSON(service.Character(c.Params("id")))
	})

	v1.Get("/enemies", func(c *fiber.Ctx) error {
		return c.JSON(service.Enemies())
	})

	v1.Get("/enemies/:id", func(c *fiber.Ctx) error {
		return c.JSON(service.Enemy(c.Params("id")))
	})

	v1.Get("/items", func(c *fiber.Ctx) error {
		return c.JSON(service.Items())
	})

	v1.Get("/items/:id", func(c *fiber.Ctx) error {
		return c.JSON(service.Item(c.Params("id")))
	})

	v1.Get("/stages", func(c *fiber.Ctx) error {
		return c.JSON(service.Stages())
	})

	v1.Get("/stages/:id", func(c *fiber.Ctx) error {
		return c.JSON(service.Stage(c.Params("id")))
	})

	v1.Get("/drops", func(c *fiber.Ctx) error {
		return c.JSON(service.Drops())
	})

	v1.Get("/drops/items", func(c *fiber.Ctx) error {
		return c.JSON(service.DropsItems())
	})

	v1.Get("/drops/stages", func(c *fiber.Ctx) error {
		return c.JSON(service.DropsStages())
	})

	v1.Get("/drops/items/:id", func(c *fiber.Ctx) error {
		return c.JSON(service.DropsItem(c.Params("id")))
	})

	v1.Get("/drops/stages/:id", func(c *fiber.Ctx) error {
		return c.JSON(service.DropsStage(c.Params("id")))
	})

	if port != nil {
		app.Listen(*port)
	} else {
		app.Listen(":3000")
	}
}
