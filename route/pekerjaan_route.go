package route

import (
	"crud-alumni/app/service"
	"crud-alumni/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterPekerjaanRoutes(api fiber.Router) {
	pekerjaan := api.Group("/pekerjaan", middleware.AuthRequired())

	pekerjaan.Get("/", service.GetPekerjaanList)
	pekerjaan.Get("/:id", service.GetPekerjaanByID)
	pekerjaan.Get("/alumni/:alumni_id", middleware.AdminOnly(), service.GetPekerjaanByAlumni)
	
	pekerjaan.Post("/", middleware.AdminOnly(), service.CreatePekerjaan)
	pekerjaan.Put("/:id", middleware.AdminOnly(), service.UpdatePekerjaan)
	pekerjaan.Delete("/:id", middleware.AdminOnly(), service.DeletePekerjaan)
}

	// pekerjaan.Get("/trash", service.TrashListHandler)
	// pekerjaan.Post("/deleted/:userid?", service.SoftDeletePekerjaan)
	// pekerjaan.Post("/restored/:userid?", service.RestorePekerjaan)
