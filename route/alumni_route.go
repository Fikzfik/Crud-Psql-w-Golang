package route

import (
	"crud-alumni/app/service"
	// "crud-alumni/helper"
	"crud-alumni/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterAlumniRoutes(api fiber.Router) {
	alumni := api.Group("/alumni", middleware.AuthRequired())

	// alumni.Get("/fakultas/:fakul", service.GetAlumniByFakultasHandler)
	alumni.Get("/", service.GetAlumniListHandler)
	alumni.Get("/:id", service.GetAlumniByIDHandler)

	alumni.Post("/", middleware.AdminOnly(), service.CreateAlumniHandler)
	alumni.Put("/:id", middleware.AdminOnly(), service.UpdateAlumniHandler)
	alumni.Delete("/:id", middleware.AdminOnly(), service.DeleteAlumniHandler)
}
