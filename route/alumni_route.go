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
	alumni.Get("/all", service.GetAlumni)
	alumni.Get("/", service.GetAlumniList)
	alumni.Get("/:id", service.GetAlumniByID)

	alumni.Post("/", middleware.AdminOnly(), service.CreateAlumni)
	alumni.Put("/:id", middleware.AdminOnly(), service.UpdateAlumni)
	alumni.Delete("/:id", middleware.AdminOnly(), service.DeleteAlumni)
}
