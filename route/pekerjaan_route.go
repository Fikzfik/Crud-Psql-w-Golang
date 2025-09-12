package route

import (
	"crud-alumni/app/models"
	"crud-alumni/app/repository"
	"crud-alumni/helper"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"crud-alumni/middleware"
)

func RegisterPekerjaanRoutes(api fiber.Router) {
	// semua route pekerjaan harus login dulu
	pekerjaan := api.Group("/pekerjaan", middleware.AuthRequired())

	pekerjaan.Get("/", func(c *fiber.Ctx) error {
		data, err := repository.GetAllPekerjaan()
		if err != nil {
			return helper.Response(c, 500, "Gagal ambil data pekerjaan", nil)
		}
		return helper.Response(c, 200, "OK", data)
	})

	pekerjaan.Get("/:id", func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		data, err := repository.GetPekerjaanByID(id)
		if err != nil {
			return helper.Response(c, 404, "Pekerjaan tidak ditemukan", nil)
		}
		return helper.Response(c, 200, "OK", data)
	})

	pekerjaan.Get("/alumni/:alumni_id", middleware.AdminOnly(), func(c *fiber.Ctx) error {
		alumniID, _ := strconv.Atoi(c.Params("alumni_id"))
		data, err := repository.GetPekerjaanByAlumni(alumniID)
		if err != nil {
			return helper.Response(c, 404, "Data pekerjaan tidak ditemukan", nil)
		}
		return helper.Response(c, 200, "OK", data)
	})

	pekerjaan.Post("/", middleware.AdminOnly(), func(c *fiber.Ctx) error {
		var p models.PekerjaanAlumni
		if err := c.BodyParser(&p); err != nil {
			return helper.Response(c, 400, "Invalid input", nil)
		}
		if err := repository.InsertPekerjaan(p); err != nil {
			return helper.Response(c, 500, "Gagal tambah pekerjaan", nil)
		}
		return helper.Response(c, 201, "Pekerjaan ditambahkan", p)
	})

	pekerjaan.Put("/:id", middleware.AdminOnly(), func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		var p models.PekerjaanAlumni
		if err := c.BodyParser(&p); err != nil {
			return helper.Response(c, 400, "Invalid input", nil)
		}
		if err := repository.UpdatePekerjaan(id, p); err != nil {
			return helper.Response(c, 500, "Gagal update pekerjaan", nil)
		}
		return helper.Response(c, 200, "Pekerjaan diupdate", p)
	})

	pekerjaan.Delete("/:id", middleware.AdminOnly(), func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		if err := repository.DeletePekerjaan(id); err != nil {
			return helper.Response(c, 500, "Gagal hapus pekerjaan", nil)
		}
		return helper.Response(c, 200, "Pekerjaan dihapus", nil)
	})
}
