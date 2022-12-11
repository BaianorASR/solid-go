package cmd

import "github.com/gofiber/fiber/v2"

func Server() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		ServerHeader:  "Solid Go",
		AppName:       "Test Solid GO App v1",
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("BAIANO")
	})

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
