package route

import (
	"crud-alumni/app/service"
	"crud-alumni/middleware"

	"github.com/gofiber/fiber/v2"
)

// RegisterFileRoutes - daftar route untuk upload file (foto & sertifikat)
func RegisterFileRoutes(api fiber.Router) {
	foto := api.Group("/foto", middleware.AuthRequired())
	foto.Post("/upload", service.UploadFoto)

	sertifikat := api.Group("/sertifikat", middleware.AuthRequired())
	sertifikat.Post("/upload", service.UploadSertifikat)
}

