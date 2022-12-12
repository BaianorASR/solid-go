package cmd

import (
	"fmt"
	"github.com/BaianorASR/solid-go/config"
	"github.com/gofiber/fiber/v2"
)

func Server() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		ServerHeader:  "Solid Go",
		AppName:       "Test Solid GO App v1",
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("BAIANO")
	})

	env := config.GetEnv()

	err := app.Listen(fmt.Sprintf(":%d", env.Port))
	if err != nil {
		panic(err)
	}
}
