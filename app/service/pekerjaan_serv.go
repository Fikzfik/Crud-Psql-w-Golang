package service

import (
	"crud-alumni/app/models"
	"crud-alumni/app/repository"
	"crud-alumni/helper"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// ===== HANDLERS =====

func GetPekerjaanListHandler(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	sortBy := c.Query("sortBy", "created_at")
	order := c.Query("order", "asc")
	search := c.Query("search", "")

	if page < 1 {
		page = 1
	}

	sortWhitelist := map[string]bool{
		"nama_perusahaan": true, "posisi_jabatan": true,
		"bidang_industri": true, "lokasi_kerja": true,
		"tanggal_mulai_kerja": true, "status_pekerjaan": true, "created_at": true,
	}
	if !sortWhitelist[sortBy] {
		sortBy = "created_at"
	}
	if strings.ToLower(order) != "desc" {
		order = "asc"
	}

	data, err := repository.GetPekerjaanWithPagination(search, sortBy, order, limit, page)
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
			Total:  int(total),
			Pages:  int((total + int64(limit) - 1) / int64(limit)),
			SortBy: sortBy,
			Order:  order,
			Search: search,
		},
	}
	return c.JSON(response)
}

func GetPekerjaanByIDHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	data, err := repository.GetPekerjaanByID(id)
	if err != nil {
		return helper.Response(c, 404, "Pekerjaan tidak ditemukan", nil)
	}
	return helper.Response(c, 200, "OK", data)
}

func GetPekerjaanByAlumniHandler(c *fiber.Ctx) error {
	alumniID := c.Params("alumni_id")
	data, err := repository.GetPekerjaanByAlumni(alumniID)
	if err != nil {
		return helper.Response(c, 404, "Data pekerjaan tidak ditemukan", nil)
	}
	return helper.Response(c, 200, "OK", data)
}

func CreatePekerjaanHandler(c *fiber.Ctx) error {
	var p models.PekerjaanAlumni
	if err := c.BodyParser(&p); err != nil {
		return helper.Response(c, 400, "Input tidak valid", nil)
	}
	if err := repository.InsertPekerjaan(p); err != nil {
		return helper.Response(c, 400, err.Error(), nil)
	}
	return helper.Response(c, 201, "Pekerjaan ditambahkan", p)
}

func UpdatePekerjaanHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	var p models.PekerjaanAlumni
	if err := c.BodyParser(&p); err != nil {
		return helper.Response(c, 400, "Input tidak valid", nil)
	}
	if err := repository.UpdatePekerjaan(id, p); err != nil {
		return helper.Response(c, 400, err.Error(), nil)
	}
	return helper.Response(c, 200, "Pekerjaan diupdate", p)
}

func DeletePekerjaanHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := repository.DeletePekerjaan(id); err != nil {
		return helper.Response(c, 500, "Gagal hapus pekerjaan", nil)
	}
	return helper.Response(c, 200, "Pekerjaan dihapus", nil)
}

// ===== LOGIKA BISNIS =====

func GetPekerjaanByID(id string) (models.PekerjaanAlumni, error) {
	return repository.GetPekerjaanByID(id)
}

func GetPekerjaanByAlumni(alumniID string) ([]models.PekerjaanAlumni, error) {
	return repository.GetPekerjaanByAlumni(alumniID)
}

func CreatePekerjaan(p models.PekerjaanAlumni) error {
	if p.NamaPerusahaan == "" || p.PosisiJabatan == "" {
		return ErrInvalidData
	}
	return repository.InsertPekerjaan(p)
}

func UpdatePekerjaan(id string, p models.PekerjaanAlumni) error {
	if p.NamaPerusahaan == "" || p.PosisiJabatan == "" {
		return ErrInvalidData
	}
	return repository.UpdatePekerjaan(id, p)
}

func DeletePekerjaan(id string) error {
	return repository.DeletePekerjaan(id)
}
