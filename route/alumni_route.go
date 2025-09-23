package route

import (
	"crud-alumni/app/models"
	"crud-alumni/app/repository"
	"crud-alumni/app/service"
	"crud-alumni/helper"
	"crud-alumni/middleware"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func RegisterAlumniRoutes(api fiber.Router) {
	alumni := api.Group("/alumni", middleware.AuthRequired())

	alumni.Get("/fakultas/:fakul", func(c *fiber.Ctx) error {
		fak := c.Params("fakul")

		fmt.Println(fak)
		var a models.Alumni
		data, err := service.GetAllAlumniByFak(fak, a)
		if err != nil {
			return helper.Response(c, 500, "Gagal ambil data alumni", nil)
		}
		return helper.Response(c, 200, "OK", data)
	})

	alumni.Get("/", func(c *fiber.Ctx) error {
		page, _ := strconv.Atoi(c.Query("page", "1"))
		limit, _ := strconv.Atoi(c.Query("limit", "10"))
		sortBy := c.Query("sortBy", "id")
		order := c.Query("order", "asc")
		search := c.Query("search", "")

		if page < 1 {
			page = 1
		}
		offset := (page - 1) * limit

		// whitelist kolom yang boleh di-sort
		sortWhitelist := map[string]bool{
			"id": true, "nim": true, "nama": true, "jurusan": true,
			"angkatan": true, "tahun_lulus": true, "email": true,
			"fakultas": true, "created_at": true,
		}
		if !sortWhitelist[sortBy] {
			sortBy = "id"
		}
		if strings.ToLower(order) != "desc" {
			order = "asc"
		}

		alumniList, err := repository.GetAlumniWithPagination(search, sortBy, order, limit, offset)
		if err != nil {
			return helper.Response(c, 500, "Gagal ambil data alumni", nil)
		}

		total, err := repository.CountAlumni(search)
		if err != nil {
			return helper.Response(c, 500, "Gagal hitung data alumni", nil)
		}

		response := models.AlumniResponse{
			Data: alumniList,
			Meta: models.MetaInfo{
				Page:   page,
				Limit:  limit,
				Total:  total,
				Pages:  (total + limit - 1) / limit,
				SortBy: sortBy,
				Order:  order,
				Search: search,
			},
		}

		return c.JSON(response)
	})

	alumni.Get("/:id", func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		data, err := service.GetAlumniByID(id)
		if err != nil {
			return helper.Response(c, 404, "Alumni tidak ditemukan", nil)
		}
		return helper.Response(c, 200, "OK", data)
	})

	alumni.Post("/", middleware.AdminOnly(), func(c *fiber.Ctx) error {
		var a models.Alumni
		if err := c.BodyParser(&a); err != nil {
			return helper.Response(c, 400, "Invalid input", nil)
		}
		if err := service.CreateAlumni(a); err != nil {
			return helper.Response(c, 500, "Gagal tambah alumni", nil)
		}
		return helper.Response(c, 201, "Alumni ditambahkan", a)
	})

	alumni.Put("/:id", middleware.AdminOnly(), func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		var a models.Alumni
		if err := c.BodyParser(&a); err != nil {
			return helper.Response(c, 400, "Invalid input", nil)
		}
		if err := service.UpdateAlumni(id, a); err != nil {
			return helper.Response(c, 500, "Gagal update alumni", nil)
		}
		return helper.Response(c, 200, "Alumni diupdate", a)
	})

	alumni.Delete("/:id", middleware.AdminOnly(), func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		if err := service.DeleteAlumni(id); err != nil {
			return helper.Response(c, 500, "Gagal hapus alumni", nil)
		}
		return helper.Response(c, 200, "Alumni dihapus", nil)
	})


}
