package helper

import "github.com/gofiber/fiber/v2"

func Response(c *fiber.Ctx, status int, message string, data interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"success": status < 400,
		"message": message,
		"data":    data,
	})
}
