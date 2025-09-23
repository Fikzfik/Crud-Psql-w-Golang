package route

import (
	"crud-alumni/app/service"
	"crud-alumni/helper"
	"crud-alumni/app/models"

	"github.com/gofiber/fiber/v2"
	"crud-alumni/middleware"
)

func RegisterAuthRoutes(api fiber.Router) {
	// public
	api.Post("/login", loginHandler)

	// protected
	protected := api.Group("", middleware.AuthRequired())
	protected.Get("/profile", profileHandler)
	protected.Get("/isdeleted", profileHandler)
}

func loginHandler(c *fiber.Ctx) error {
	var req models.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return helper.Response(c, 400, "Request body tidak valid", nil)
	}

	resp, err := service.Login(req)
	if err != nil {
		return helper.Response(c, 401, err.Error(), nil)
	}

	return helper.Response(c, 200, "Login berhasil", resp)
}

func profileHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(int)
	username := c.Locals("username").(string)
	role := c.Locals("role").(string)

	return helper.Response(c, 200, "Profile berhasil diambil", fiber.Map{
		"user_id":  userID,
		"username": username,
		"role":     role,
	})
}
