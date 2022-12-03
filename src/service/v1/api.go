package apiv1

import (
	"prts-backend/src/model"

	"github.com/gofiber/fiber/v2"
)

func GetAPI() *fiber.App {
	api := fiber.New()

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("prts-backend api version 1 is running!")
	})

	//* RESTAPI>V1>CHARACTER

	//* RESTAPI>V1>CHARACTER>GET
	api.Get("/characters", func(c *fiber.Ctx) error {
		return c.JSON(Characters())
	})

	api.Get("/characters/:id", func(c *fiber.Ctx) error {
		return c.JSON(Character(c.Params("id")))
	})

	//* RESTAPI>V1>ENEMY

	//* RESTAPI>V1>ENEMY>GET
	api.Get("/enemies", func(c *fiber.Ctx) error {
		return c.JSON(Enemies())
	})

	api.Get("/enemies/:id", func(c *fiber.Ctx) error {
		return c.JSON(Enemy(c.Params("id")))
	})

	//* RESTAPI>V1>ITEM

	//* RESTAPI>V1>ITEM>GET
	api.Get("/items", func(c *fiber.Ctx) error {
		return c.JSON(Items())
	})

	api.Get("/items/:id", func(c *fiber.Ctx) error {
		return c.JSON(Item(c.Params("id")))
	})

	//* RESTAPI>V1>STAGE

	//* RESTAPI>V1>STAGE>GET
	api.Get("/stages", func(c *fiber.Ctx) error {
		return c.JSON(Stages())
	})

	api.Get("/stages/:id", func(c *fiber.Ctx) error {
		return c.JSON(Stage(c.Params("id")))
	})

	//* RESTAPI>V1>DROP

	//* RESTAPI>V1>DROP>GET
	api.Get("/drops", func(c *fiber.Ctx) error {
		result, err := Drops()
		if err != nil {
			return err
		}
		return c.JSON(result)
	})

	api.Get("/drops/items", func(c *fiber.Ctx) error {
		result, err := DropsItems()
		if err != nil {
			return err
		}
		return c.JSON(result)
	})

	api.Get("/drops/stages", func(c *fiber.Ctx) error {
		result, err := DropsStages()
		if err != nil {
			return err
		}
		return c.JSON(result)
	})

	api.Get("/drops/items/:id", func(c *fiber.Ctx) error {
		return c.JSON(DropsItem(c.Params("id")))
	})

	api.Get("/drops/stages/:id", func(c *fiber.Ctx) error {
		return c.JSON(DropsStage(c.Params("id")))
	})

	//* RESTAPI>V1>DROP>POST
	api.Post("/drop", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		var drop model.Drop
		c.BodyParser(&drop)
		return CreateDrop(&drop)
	})

	api.Post("/drops", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		var drops []model.Drop
		c.BodyParser(&drops)
		return CreateDrops(drops)
	})

	//* RESTAPI>V1>DROP>PUT
	api.Put("/drop", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		var drop model.Drop
		c.BodyParser(&drop)
		return UpdateDrop(&drop)
	})

	api.Put("/drops", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		var drops []model.Drop
		c.BodyParser(&drops)
		return UpdateDrops(drops)
	})

	//* RESTAPI>V1>DROP>DELETE
	api.Delete("/drop", func(c *fiber.Ctx) error {
		//! DELETE A DROP IS NOT ALLOWED
		return fiber.NewError(403, "PERMISSION DENIED: DELETE DROP")
	})

	api.Delete("/drops", func(c *fiber.Ctx) error {
		//! DELETE DROPS IS NOT ALLOWED
		return fiber.NewError(403, "PERMISSION DENIED: DELETE DROPS")
	})

	return api
}
