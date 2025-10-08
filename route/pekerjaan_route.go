package route

import (
	"crud-alumni/app/models"
	"crud-alumni/app/repository"
	"crud-alumni/app/service"
	"crud-alumni/helper"
	"crud-alumni/middleware"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func RegisterPekerjaanRoutes(api fiber.Router) {
	// semua route pekerjaan harus login dulu
	pekerjaan := api.Group("/pekerjaan", middleware.AuthRequired())

	pekerjaan.Get("/isdeleted", func(c *fiber.Ctx) error {
		role := c.Locals("role").(string)
		userID := c.Locals("user_id").(int)

		if err := service.IsDeleted(role, userID); err != nil {
			return helper.Response(c, 500, "Gagal hapus alumni", nil)
		}
		return helper.Response(c, 200, "Alumni dihapus", nil)
	})
	pekerjaan.Get("/aktif", func(c *fiber.Ctx) error {
		data, err := service.IsAktif()
		if err != nil {
			return helper.Response(c, 404, "Pekerjaan tidak ditemukans", data)
		}
		return helper.Response(c, 200, "OK", data)
	})

	
	pekerjaan.Get("/isrestored", func(c *fiber.Ctx) error {
		role := c.Locals("role").(string)
		userID := c.Locals("user_id").(int)
		
		if err := service.IsRestored(role, userID); err != nil {
			return helper.Response(c, 500, "Gagal hapus alumni", nil)
		}
		return helper.Response(c, 200, "Alumni dikembalikan", nil)
	})
	pekerjaan.Get("/:id", func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		data, err := repository.GetPekerjaanByID(id)
		if err != nil {
			return helper.Response(c, 404, "Pekerjaan tidak ditemukan", nil)
		}
		return helper.Response(c, 200, "OK", data)
	})

	pekerjaan.Post("/restored/:userid", middleware.AdminOnly(), func(c *fiber.Ctx) error {
		alumniID, _ := strconv.Atoi(c.Params("userid"))
		var p models.PekerjaanAlumni
		if err := c.BodyParser(&p); err != nil {
			return helper.Response(c, 400, "Invalid input", nil)
		}
		data, err := service.RestoredUser(alumniID, p)
		if err != nil {
			return helper.Response(c, 500, "berhasil ubah user lain", nil)
		}
		return helper.Response(c, 200, "OK", data)
	})

	pekerjaan.Get("/", func(c *fiber.Ctx) error {
		page, _ := strconv.Atoi(c.Query("page", "1"))
		limit, _ := strconv.Atoi(c.Query("limit", "10"))
		sortBy := c.Query("sortBy", "id")
		order := c.Query("order", "asc")
		search := c.Query("search", "")

		if page < 1 {
			page = 1
		}
		offset := (page - 1) * limit

		// whitelist kolom yang bisa disort
		sortWhitelist := map[string]bool{
			"id": true, "nama_perusahaan": true, "posisi_jabatan": true,
			"bidang_industri": true, "lokasi_kerja": true,
			"tanggal_mulai_kerja": true, "status_pekerjaan": true, "created_at": true,
		}
		if !sortWhitelist[sortBy] {
			sortBy = "id"
		}
		if strings.ToLower(order) != "desc" {
			order = "asc"
		}

		data, err := repository.GetPekerjaanWithPagination(search, sortBy, order, limit, offset)
		if err != nil {
			return helper.Response(c, 500, "Gagal ambil data pekerjaan", nil)
		}

		total, err := repository.CountPekerjaan(search)
		if err != nil {
			return helper.Response(c, 500, "Gagal hitung data pekerjaan", nil)
		}

		response := models.PekerjaanResponse{
			Data: data,
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
