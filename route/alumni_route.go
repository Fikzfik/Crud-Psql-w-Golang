package route

import (
	"crud-alumni/helper"
	"crud-alumni/model"
	"crud-alumni/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func RegisterAlumniRoutes(api fiber.Router) {
	alumni := api.Group("/alumni")

	alumni.Get("/", func(c *fiber.Ctx) error {
		data, err := service.GetAllAlumni()
		if err != nil {
			return helper.Response(c, 500, "Gagal ambil data alumni", nil)
		}
		return helper.Response(c, 200, "OK", data)
	})

	alumni.Get("/:id", func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		data, err := service.GetAlumniByID(id)
		if err != nil {
			return helper.Response(c, 404, "Alumni tidak ditemukan", nil)
		}
		return helper.Response(c, 200, "OK", data)
	})

	alumni.Post("/", func(c *fiber.Ctx) error {
		var a model.Alumni
		if err := c.BodyParser(&a); err != nil {
			return helper.Response(c, 400, "Invalid input", nil)
		}
		if err := service.CreateAlumni(a); err != nil {
			return helper.Response(c, 500, "Gagal tambah alumni", nil)
		}
		return helper.Response(c, 201, "Alumni ditambahkan", a)
	})

	alumni.Put("/:id", func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		var a model.Alumni
		if err := c.BodyParser(&a); err != nil {
			return helper.Response(c, 400, "Invalid input", nil)
		}
		if err := service.UpdateAlumni(id, a); err != nil {
			return helper.Response(c, 500, "Gagal update alumni", nil)
		}
		return helper.Response(c, 200, "Alumni diupdate", a)
	})

	alumni.Delete("/:id", func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		if err := service.DeleteAlumni(id); err != nil {
			return helper.Response(c, 500, "Gagal hapus alumni", nil)
		}
		return helper.Response(c, 200, "Alumni dihapus", nil)
	})
}
