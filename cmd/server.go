package cmd

import (
	"fmt"

	"github.com/BaianorASR/solid-go/config/environment"
	"github.com/gofiber/fiber/v2"
)

func Server() {
	env := config.GetEnv()

	app := fiber.New(fiber.Config{
		CaseSensitive: env.CaseSensitiveURL,
		ServerHeader:  "Solid Go",
		AppName:       "Test Solid GO App v1",
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("BAIANO")
	})

	err := app.Listen(fmt.Sprintf(":%d", env.Port))
	if err != nil {
		panic(err)
	}
}
