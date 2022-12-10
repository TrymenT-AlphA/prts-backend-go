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

	//********************* CHARACTER ***************************

	api.Get("/characters", func(c *fiber.Ctx) error {
		return c.JSON(Characters())
	})

	api.Get("/characters/:id", func(c *fiber.Ctx) error {
		return c.JSON(Character(c.Params("id")))
	})

	//********************* ENEMY ***************************

	api.Get("/enemies", func(c *fiber.Ctx) error {
		return c.JSON(Enemies())
	})

	api.Get("/enemies/:id", func(c *fiber.Ctx) error {
		return c.JSON(Enemy(c.Params("id")))
	})

	//********************* ITEM ***************************

	api.Get("/items", func(c *fiber.Ctx) error {
		return c.JSON(Items())
	})

	api.Get("/items/:id", func(c *fiber.Ctx) error {
		return c.JSON(Item(c.Params("id")))
	})

	//********************* STAGE ***************************

	api.Get("/stages", func(c *fiber.Ctx) error {
		return c.JSON(Stages())
	})

	api.Get("/stages/:id", func(c *fiber.Ctx) error {
		return c.JSON(Stage(c.Params("id")))
	})

	//********************* DROP ***************************

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

	api.Post("/drop", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		var drop model.Drop
		if err := c.BodyParser(&drop); err != nil {
			return err
		}
		return CreateDrop(&drop)
	})

	api.Post("/drops", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		var drops []model.Drop
		if err := c.BodyParser(&drops); err != nil {
			return err
		}
		return CreateDrops(drops)
	})

	api.Put("/drop", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		var drop model.Drop
		if err := c.BodyParser(&drop); err != nil {
			return err
		}
		return UpdateDrop(&drop)
	})

	api.Put("/drops", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		var drops []model.Drop
		if err := c.BodyParser(&drops); err != nil {
			return err
		}
		return UpdateDrops(drops)
	})

	api.Delete("/drop", func(c *fiber.Ctx) error {
		return fiber.NewError(403, "PERMISSION DENIED: DELETE DROP")
	})

	api.Delete("/drops", func(c *fiber.Ctx) error {
		return fiber.NewError(403, "PERMISSION DENIED: DELETE DROPS")
	})

	return api
}
