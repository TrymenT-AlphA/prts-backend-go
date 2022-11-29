package server

import (
	"log"
	"prts-backend/src/model"
	"prts-backend/src/service/v1"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Run(port *string, config *fiber.Config) {

	if !fiber.IsChild() {
		err := service.InitDatabase(true)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := service.InitDatabase(false)
		if err != nil {
			log.Fatal(err)
		}
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

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("prts-backend service is running!")
	})

	//* RESTAPI
	api := app.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("prts-backend api is running!")
	})

	//* RESTAPI>V1
	v1 := api.Group("/v1")

	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("prts-backend api version 1 is running!")
	})

	//* RESTAPI>V1>CHARACTER

	//* RESTAPI>V1>CHARACTER>GET
	v1.Get("/characters", func(c *fiber.Ctx) error {
		return c.JSON(service.Characters())
	})

	v1.Get("/characters/:id", func(c *fiber.Ctx) error {
		return c.JSON(service.Character(c.Params("id")))
	})

	//* RESTAPI>V1>ENEMY

	//* RESTAPI>V1>ENEMY>GET
	v1.Get("/enemies", func(c *fiber.Ctx) error {
		return c.JSON(service.Enemies())
	})

	v1.Get("/enemies/:id", func(c *fiber.Ctx) error {
		return c.JSON(service.Enemy(c.Params("id")))
	})

	//* RESTAPI>V1>ITEM

	//* RESTAPI>V1>ITEM>GET
	v1.Get("/items", func(c *fiber.Ctx) error {
		return c.JSON(service.Items())
	})

	v1.Get("/items/:id", func(c *fiber.Ctx) error {
		return c.JSON(service.Item(c.Params("id")))
	})

	//* RESTAPI>V1>STAGE

	//* RESTAPI>V1>STAGE>GET
	v1.Get("/stages", func(c *fiber.Ctx) error {
		return c.JSON(service.Stages())
	})

	v1.Get("/stages/:id", func(c *fiber.Ctx) error {
		return c.JSON(service.Stage(c.Params("id")))
	})

	//* RESTAPI>V1>DROP

	//* RESTAPI>V1>DROP>GET
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

	//* RESTAPI>V1>DROP>POST
	v1.Post("/drop", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		var drop model.Drop
		c.BodyParser(&drop)
		return service.CreateDrop(&drop)
	})

	v1.Post("/drops", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		var drops []model.Drop
		c.BodyParser(&drops)
		return service.CreateDrops(drops)
	})

	//* RESTAPI>V1>DROP>PUT
	v1.Put("/drop", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		var drop model.Drop
		c.BodyParser(&drop)
		return service.UpdateDrop(&drop)
	})

	v1.Put("/drops", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		var drops []model.Drop
		c.BodyParser(&drops)
		return service.UpdateDrops(drops)
	})

	//* RESTAPI>V1>DROP>DELETE
	v1.Delete("/drop", func(c *fiber.Ctx) error {
		//! DELETE A DROP IS NOT ALLOWED
		return fiber.NewError(403, "DANGER ACTION: DELETE DROP")
	})

	v1.Delete("/drops", func(c *fiber.Ctx) error {
		//! DELETE DROPS IS NOT ALLOWED
		return fiber.NewError(403, "DANGER ACTION: DELETE DROPS")
	})

	if port != nil {
		err := app.Listen(*port)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := app.Listen(":3000")
		if err != nil {
			log.Fatal(err)
		}
	}

}
