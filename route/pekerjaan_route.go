package route

import (
	"crud-alumni/app/service"
	"crud-alumni/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterPekerjaanRoutes(api fiber.Router) {
	pekerjaan := api.Group("/pekerjaan", middleware.AuthRequired())

	pekerjaan.Get("/", service.GetPekerjaanListHandler)
	pekerjaan.Get("/:id", service.GetPekerjaanByIDHandler)
	pekerjaan.Get("/alumni/:alumni_id", middleware.AdminOnly(), service.GetPekerjaanByAlumniHandler)
	
	// pekerjaan.Get("/trash", service.TrashListHandler)
	// pekerjaan.Post("/deleted/:userid?", service.SoftDeletePekerjaan)
	// pekerjaan.Post("/restored/:userid?", service.RestorePekerjaan)

	pekerjaan.Post("/", middleware.AdminOnly(), service.CreatePekerjaanHandler)
	pekerjaan.Put("/:id", middleware.AdminOnly(), service.UpdatePekerjaanHandler)
	pekerjaan.Delete("/:id", middleware.AdminOnly(), service.DeletePekerjaanHandler)
}
