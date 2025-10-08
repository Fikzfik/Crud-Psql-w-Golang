package route

import (
	"crud-alumni/app/service"
	"crud-alumni/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterPekerjaanRoutes(api fiber.Router) {
	pekerjaan := api.Group("/pekerjaan", middleware.AuthRequired())

	pekerjaan.Get("/deleted", service.SoftDeleteHandler)
	pekerjaan.Get("/trash", service.TrashListHandler)
	pekerjaan.Get("/restored", service.RestoreSelfHandler)
	pekerjaan.Get("/", service.GetPekerjaanListHandler)
	pekerjaan.Get("/:id", service.GetPekerjaanByIDHandler)
	pekerjaan.Get("/alumni/:alumni_id", middleware.AdminOnly(), service.GetPekerjaanByAlumniHandler)


	pekerjaan.Post("/deleted/:userid", middleware.AdminOnly(), service.DeleteUserPekerjaanHandler)
	pekerjaan.Post("/restored/:userid", middleware.AdminOnly(), service.RestoreUserPekerjaanHandler)

	pekerjaan.Post("/", middleware.AdminOnly(), service.CreatePekerjaanHandler)
	pekerjaan.Put("/:id", middleware.AdminOnly(), service.UpdatePekerjaanHandler)
	pekerjaan.Delete("/:id", middleware.AdminOnly(), service.DeletePekerjaanHandler)
}
