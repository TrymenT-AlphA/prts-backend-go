# prts-backend
prts-backend is a REST style api using golang

## ADD YOUR ENV

```.env
/.env
<user>:<pwd>@tcp(<host>:<port>)/prts-go?charset=utf8mb4&parseTime=True&loc=Local
```

## BUILD && RUN (USING DOCKER RECOMANDED)

```cmd
docker compose --build up
```

## BUILD && RUN

```cmd
./cmd/run.bat
```

## API

see /src/server/run.go > func Run
```go
func Run(port *string, config *fiber.Config) {

    ...

	app.Use(cors.New())

	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("prts-backend service is running!")
	})

	api := app.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("prts-backend api is running!")
	})

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

    ...

}
```