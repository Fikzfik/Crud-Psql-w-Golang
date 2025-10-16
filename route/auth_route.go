package route

import (
	"crud-alumni/app/service"
	"crud-alumni/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(api fiber.Router) {
	api.Post("/login", service.LoginHandler)

	protected := api.Group("", middleware.AuthRequired())
	protected.Get("/profile", service.ProfileHandler)
}
