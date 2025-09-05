package route

import (
	"crud-alumni/helper"
	"crud-alumni/model"
	"crud-alumni/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func RegisterPekerjaanRoutes(api fiber.Router) {
	pekerjaan := api.Group("/pekerjaan")

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

	pekerjaan.Get("/alumni/:alumni_id", func(c *fiber.Ctx) error {
		alumniID, _ := strconv.Atoi(c.Params("alumni_id"))
		data, err := repository.GetPekerjaanByAlumni(alumniID)
		if err != nil {
			return helper.Response(c, 404, "Data pekerjaan tidak ditemukan", nil)
		}
		return helper.Response(c, 200, "OK", data)
	})

	pekerjaan.Post("/", func(c *fiber.Ctx) error {
		var p model.PekerjaanAlumni
		if err := c.BodyParser(&p); err != nil {
			return helper.Response(c, 400, "Invalid input", nil)
		}
		if err := repository.InsertPekerjaan(p); err != nil {
			return helper.Response(c, 500, "Gagal tambah pekerjaan", nil)
		}
		return helper.Response(c, 201, "Pekerjaan ditambahkan", p)
	})

	pekerjaan.Put("/:id", func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		var p model.PekerjaanAlumni
		if err := c.BodyParser(&p); err != nil {
			return helper.Response(c, 400, "Invalid input", nil)
		}
		if err := repository.UpdatePekerjaan(id, p); err != nil {
			return helper.Response(c, 500, "Gagal update pekerjaan", nil)
		}
		return helper.Response(c, 200, "Pekerjaan diupdate", p)
	})

	pekerjaan.Delete("/:id", func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		if err := repository.DeletePekerjaan(id); err != nil {
			return helper.Response(c, 500, "Gagal hapus pekerjaan", nil)
		}
		return helper.Response(c, 200, "Pekerjaan dihapus", nil)
	})
}
