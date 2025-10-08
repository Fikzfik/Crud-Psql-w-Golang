package service

import (
	"crud-alumni/app/models"
	"crud-alumni/app/repository"
	"crud-alumni/helper"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// ===== HANDLER =====

func GetPekerjaanListHandler(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	sortBy := c.Query("sortBy", "id")
	order := c.Query("order", "asc")
	search := c.Query("search", "")

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * limit

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
}

func GetPekerjaanByIDHandler(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	data, err := GetPekerjaanByID(id)
	if err != nil {
		return helper.Response(c, 404, "Pekerjaan tidak ditemukanss", nil)
	}
	return helper.Response(c, 200, "OK", data)
}

func GetPekerjaanByAlumniHandler(c *fiber.Ctx) error {
	alumniID, _ := strconv.Atoi(c.Params("alumni_id"))
	data, err := GetPekerjaanByAlumni(alumniID)
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
	if err := CreatePekerjaan(p); err != nil {
		return helper.Response(c, 400, err.Error(), nil)
	}
	return helper.Response(c, 201, "Pekerjaan ditambahkan", p)
}

func UpdatePekerjaanHandler(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var p models.PekerjaanAlumni
	if err := c.BodyParser(&p); err != nil {
		return helper.Response(c, 400, "Input tidak valid", nil)
	}
	if err := UpdatePekerjaan(id, p); err != nil {
		return helper.Response(c, 400, err.Error(), nil)
	}
	return helper.Response(c, 200, "Pekerjaan diupdate", p)
}

func DeletePekerjaanHandler(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := DeletePekerjaan(id); err != nil {
		return helper.Response(c, 500, "Gagal hapus pekerjaan", nil)
	}
	return helper.Response(c, 200, "Pekerjaan dihapus", nil)
}

// Soft delete / restore handlers
func SoftDeleteHandler(c *fiber.Ctx) error {
	role := c.Locals("role").(string)
	id := c.Locals("user_id").(int)
	if err := IsDeleted(role, id); err != nil {
		return helper.Response(c, 500, "Gagal hapus pekerjaan", nil)
	}
	return helper.Response(c, 200, "Data dihapus", nil)
}

func TrashListHandler(c *fiber.Ctx) error {
	role := c.Locals("role").(string)
	id := c.Locals("user_id").(int)
	data, err := IsAktif(role, id)
	if err != nil {
		return helper.Response(c, 404, "Data tidak ditemukan", nil)
	}
	return helper.Response(c, 200, "Data trash", data)
}

func RestoreSelfHandler(c *fiber.Ctx) error {
	role := c.Locals("role").(string)
	id := c.Locals("user_id").(int)
	if err := IsRestored(role, id); err != nil {
		return helper.Response(c, 500, "Gagal restore data", nil)
	}
	return helper.Response(c, 200, "Data berhasil direstore", nil)
}

func DeleteUserPekerjaanHandler(c *fiber.Ctx) error {
	alumniID, _ := strconv.Atoi(c.Params("userid"))
	var p models.PekerjaanAlumni
	if err := c.BodyParser(&p); err != nil {
		return helper.Response(c, 400, "Input tidak valid", nil)
	}
	data, err := Userdeleted(alumniID, p)
	if err != nil {
		return helper.Response(c, 500, "Gagal hapus data user lain", nil)
	}
	return helper.Response(c, 200, "User dihapus", data)
}

func RestoreUserPekerjaanHandler(c *fiber.Ctx) error {
	alumniID, _ := strconv.Atoi(c.Params("userid"))
	var p models.PekerjaanAlumni
	if err := c.BodyParser(&p); err != nil {
		return helper.Response(c, 400, "Input tidak valid", nil)
	}
	data, err := RestoredUser(alumniID, p)
	if err != nil {
		return helper.Response(c, 500, "Gagal restore data user lain", nil)
	}
	return helper.Response(c, 200, "User direstore", data)
}

// ===== LOGIKA BISNIS =====

func GetPekerjaanByID(id int) (models.PekerjaanAlumni, error) {
	return repository.GetPekerjaanByID(id)
}

func GetPekerjaanByAlumni(alumniID int) ([]models.PekerjaanAlumni, error) {
	return repository.GetPekerjaanByAlumni(alumniID)
}

func CreatePekerjaan(p models.PekerjaanAlumni) error {
	if p.NamaPerusahaan == "" || p.PosisiJabatan == "" {
		return ErrInvalidData
	}
	return repository.InsertPekerjaan(p)
}

func UpdatePekerjaan(id int, p models.PekerjaanAlumni) error {
	if p.NamaPerusahaan == "" || p.PosisiJabatan == "" {
		return ErrInvalidData
	}
	return repository.UpdatePekerjaan(id, p)
}

func DeletePekerjaan(id int) error {
	return repository.DeletePekerjaan(id)
}

func IsDeleted(role string, id int) error {
	return repository.IsDeleted(role, id)
}

func IsAktif(role string, id int) ([]models.PekerjaanAlumni, error) {
	return repository.IsAktif(role, id)
}

func IsRestored(role string, id int) error {
	return repository.IsRestored(role, id)
}

func RestoredUser(id int, p models.PekerjaanAlumni) ([]models.PekerjaanAlumni, error) {
	return repository.RestoredUser(id, p)
}

func Userdeleted(id int, p models.PekerjaanAlumni) ([]models.PekerjaanAlumni, error) {
	return repository.Userdeleted(id, p)
}
