package service

import (
	"crud-alumni/app/models"
	"crud-alumni/app/repository"
	"crud-alumni/helper"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ===== HANDLERS =====
// DeletePekerjaan godoc
// @Summary Menghapus data pekerjaan
// @Description Menghapus pekerjaan alumni berdasarkan ID
// @Tags Pekerjaan
// @Accept json
// @Produce json
// @Param id path string true "ID Pekerjaan"
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /pekerjaan/{id} [delete]
func GetPekerjaanList(c *fiber.Ctx) error {
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

// GetPekerjaanByID godoc
// @Summary Mendapatkan data pekerjaan berdasarkan ID
// @Description Mengambil detail pekerjaan berdasarkan ID yang diberikan
// @Tags Pekerjaan
// @Accept json
// @Produce json
// @Param id path string true "ID Pekerjaan"
// @Security BearerAuth
// @Success 200 {object} models.PekerjaanAlumni
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /pekerjaan/{id} [get]

func GetPekerjaanByID(c *fiber.Ctx) error {
	id := c.Params("id")
	data, err := repository.GetPekerjaanByID(id)
	if err != nil {
		return helper.Response(c, 404, "Pekerjaan tidak ditemukan", nil)
	}
	return helper.Response(c, 200, "OK", data)
}

// GetAllPekerjaan godoc
// @Summary      Mendapatkan semua data pekerjaan
// @Description  Mengambil seluruh data pekerjaan dari database tanpa pagination atau filter
// @Tags         Pekerjaan
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   models.PekerjaanAlumni  "Daftar seluruh Pekerjaan Alumni"
// @Failure      404  {object}  map[string]interface{} "Data pekerjaan tidak ditemukan"
// @Router       /pekerjaan/all [get]
func GetAllPekerjaan(c *fiber.Ctx) error {
	id := c.Params("id")
	data, err := repository.GetPekerjaanByID(id)
	if err != nil {
		return helper.Response(c, 404, "Pekerjaan tidak ditemukan", nil)
	}
	return helper.Response(c, 200, "OK", data)
}
// GetPekerjaanByAlumni godoc
// @Summary Mendapatkan data pekerjaan berdasarkan ID alumni
// @Description Mengambil semua pekerjaan yang dimiliki oleh satu alumni tertentu
// @Tags Pekerjaan
// @Accept json
// @Produce json
// @Param alumni_id path string true "ID Alumni"
// @Security BearerAuth
// @Success 200 {array} models.PekerjaanAlumni
// @Failure 404 {object} map[string]interface{}
// @Router /pekerjaan/alumni/{alumni_id} [get]
func GetPekerjaanByAlumni(c *fiber.Ctx) error {
	alumniID := c.Params("alumni_id")
	data, err := repository.GetPekerjaanByAlumni(alumniID)
	if err != nil {
		return helper.Response(c, 404, "Data pekerjaan tidak ditemukan", nil)
	}
	return helper.Response(c, 200, "OK", data)
}

// CreatePekerjaan godoc
// @Summary Menambahkan data pekerjaan baru
// @Description Menambahkan data pekerjaan alumni baru ke database
// @Tags Pekerjaan
// @Accept json
// @Produce json
// @Param pekerjaan body models.PekerjaanAlumni true "Data pekerjaan"
// @Security BearerAuth
// @Success 201 {object} models.PekerjaanAlumni
// @Failure 400 {object} map[string]interface{}
// @Router /pekerjaan [post]
func CreatePekerjaan(c *fiber.Ctx) error {
	var p models.PekerjaanAlumni
	if err := c.BodyParser(&p); err != nil {
		return helper.Response(c, 400, "Input tidak valid", nil)
	}

	// Convert alumni_id string → ObjectID
	alumniID, err := primitive.ObjectIDFromHex(p.AlumniID.Hex())
	if err != nil {
		return helper.Response(c, 400, "alumni_id tidak valid", nil)
	}
	p.AlumniID = alumniID

	if err := repository.InsertPekerjaan(p); err != nil {
		return helper.Response(c, 400, err.Error(), nil)
	}

	return helper.Response(c, 201, "Pekerjaan ditambahkan", p)
}


// UpdatePekerjaan godoc
// @Summary Update data pekerjaan alumni
// @Description Mengupdate data pekerjaan berdasarkan ID
// @Tags Pekerjaan
// @Accept json
// @Produce json
// @Param id path string true "ID Pekerjaan"
// @Param pekerjaan body models.PekerjaanAlumni true "Data pekerjaan"
// @Security BearerAuth
// @Success 200 {object} models.PekerjaanAlumni
// @Failure 400 {object} map[string]interface{}
// @Router /pekerjaan/{id} [put]
func UpdatePekerjaan(c *fiber.Ctx) error {
	id := c.Params("id")
	var p models.PekerjaanAlumni

	if err := c.BodyParser(&p); err != nil {
		return helper.Response(c, 400, "Input tidak valid", nil)
	}

	// 🔹 Update dan ambil ulang data dari DB
	updated, err := repository.UpdatePekerjaan(id, p)
	if err != nil {
		return helper.Response(c, 400, err.Error(), nil)
	}

	// 🔹 Kembalikan hasilnya
	return helper.Response(c, 200, "Pekerjaan diupdate", updated)
}

// DeletePekerjaan godoc
// @Summary Menghapus data pekerjaan
// @Description Menghapus pekerjaan alumni berdasarkan ID
// @Tags Pekerjaan
// @Accept json
// @Produce json
// @Param id path string true "ID Pekerjaan"
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /pekerjaan/{id} [delete]
func DeletePekerjaan(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := repository.DeletePekerjaan(id); err != nil {
		return helper.Response(c, 500, "Gagal hapus pekerjaan", nil)
	}
	return helper.Response(c, 200, "Pekerjaan dihapus", nil)
}

// ===== LOGIKA BISNIS =====

func GetPekerjaanByIDs(id string) (models.PekerjaanAlumni, error) {
	return repository.GetPekerjaanByID(id)
}

func GetPekerjaanByAlumnis(alumniID string) ([]models.PekerjaanAlumni, error) {
	return repository.GetPekerjaanByAlumni(alumniID)
}

func CreatePekerjaans(p models.PekerjaanAlumni) error {
	if p.NamaPerusahaan == "" || p.PosisiJabatan == "" {
		return ErrInvalidData
	}
	return repository.InsertPekerjaan(p)
}

// func UpdatePekerjaans(id string, p models.PekerjaanAlumni) error {
// 	if p.NamaPerusahaan == "" || p.PosisiJabatan == "" {
// 		return ErrInvalidData
// 	}
// 	return repository.UpdatePekerjaan(id, p)
// }

func DeletePekerjaans(id string) error {
	return repository.DeletePekerjaan(id)
}
