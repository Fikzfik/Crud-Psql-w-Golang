package config

import "github.com/gofiber/fiber/v2"

func NewApp() *fiber.App {
	app := fiber.New()
	return app
}
