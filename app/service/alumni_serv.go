package service

import (
	"crud-alumni/app/models"
	"crud-alumni/app/repository"
	"crud-alumni/helper"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"strings"
)

// ====== HANDLER ======
// GetAlumniList godoc
// @Summary      Mendapatkan daftar alumni
// @Description  Mengambil data alumni dengan pagination, search, dan sorting
// @Tags         Alumni
// @Accept       json
// @Produce      json
// @Param        page     query     int     false  "Halaman saat ini (default 1)"
// @Param        limit    query     int     false  "Jumlah data per halaman (default 10)"
// @Param        sortBy   query     string  false  "Kolom pengurutan (default: created_at)"
// @Param        order    query     string  false  "Arah pengurutan (asc/desc)"
// @Param        search   query     string  false  "Kata kunci pencarian"
// @Security     BearerAuth
// @Success      200  {object}  models.AlumniResponse
// @Failure      500  {object}  map[string]interface{}
// @Router       /alumni [get]
func GetAlumniList(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	sortBy := c.Query("sortBy", "created_at")
	order := c.Query("order", "asc")
	search := c.Query("search", "")

	if page < 1 {
		page = 1
	}

	sortWhitelist := map[string]bool{
		"nama": true, "jurusan": true, "angkatan": true,
		"tahun_lulus": true, "fakultas": true, "created_at": true,
	}
	if !sortWhitelist[sortBy] {
		sortBy = "created_at"
	}
	if strings.ToLower(order) != "desc" {
		order = "asc"
	}

	alumniList, err := repository.GetAlumniWithPagination(search, sortBy, order, limit, page)
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
			Total:  int(total),
			Pages:  int((total + int64(limit) - 1) / int64(limit)),
			SortBy: sortBy,
			Order:  order,
			Search: search,
		},
	}

	return c.Status(200).JSON(response)
}

// GetAllAlumni godoc
// @Summary      Mendapatkan semua data alumni
// @Description  Mengambil seluruh data alumni dari database tanpa pagination atau filter
// @Tags         Alumni
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   models.Alumni  "Daftar seluruh alumni"
// @Failure      404  {object}  map[string]interface{} "Data alumni tidak ditemukan"
// @Router       /alumni/all [get]
func GetAlumni(c *fiber.Ctx) error {
	data, err := repository.GetAllAlumni()
	if err != nil {
		return helper.Response(c, 404, "Alumni tidak ditemukan", nil)
	}
	return helper.Response(c, 200, "OK", data)
}

// GetAlumniByID godoc
// @Summary      Mendapatkan detail alumni berdasarkan ID
// @Description  Mengambil satu data alumni berdasarkan ID
// @Tags         Alumni
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "ID Alumni"
// @Security     BearerAuth
// @Success      200  {object}  models.Alumni
// @Failure      404  {object}  map[string]interface{}
// @Router       /alumni/{id} [get]
func GetAlumniByID(c *fiber.Ctx) error {
	id := c.Params("id")
	data, err := repository.GetAlumniByID(id)
	if err != nil {
		return helper.Response(c, 404, "Alumni tidak ditemukan", nil)
	}
	return helper.Response(c, 200, "OK", data)
}


// CreateAlumni godoc
// @Summary      Menambahkan alumni baru
// @Description  Membuat data alumni baru (hanya admin)
// @Tags         Alumni
// @Accept       json
// @Produce      json
// @Param        alumni  body  models.Alumni  true  "Data Alumni"
// @Security     BearerAuth
// @Success      201  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Router       /alumni [post]
func CreateAlumni(c *fiber.Ctx) error {
	var a models.Alumni
	if err := c.BodyParser(&a); err != nil {
		return helper.Response(c, 400, "Input tidak valid", nil)
	}

	if err := repository.InsertAlumni(a); err != nil {
		return helper.Response(c, 400, err.Error(), nil)
	}

	return helper.Response(c, 201, "Alumni ditambahkan", a)
}

// UpdateAlumni godoc
// @Summary      Mengupdate data alumni
// @Description  Update data alumni berdasarkan ID (hanya admin)
// @Tags         Alumni
// @Accept       json
// @Produce      json
// @Param        id      path  string         true  "ID Alumni"
// @Param        alumni  body  models.Alumni  true  "Data Alumni yang diupdate"
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Router       /alumni/{id} [put]
func UpdateAlumni(c *fiber.Ctx) error {
	id := c.Params("id")
	var a models.Alumni
	if err := c.BodyParser(&a); err != nil {
		return helper.Response(c, 400, "Input tidak valid", nil)
	}
	if err := repository.UpdateAlumni(id, a); err != nil {
		return helper.Response(c, 400, err.Error(), nil)
	}
	return helper.Response(c, 200, "Alumni diupdate", a)
}

// DeleteAlumni godoc
// @Summary      Menghapus alumni
// @Description  Menghapus data alumni berdasarkan ID (hanya admin)
// @Tags         Alumni
// @Accept       json
// @Produce      json
// @Param        id   path  string  true  "ID Alumni"
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /alumni/{id} [delete]
func DeleteAlumni(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := repository.DeleteAlumni(id); err != nil {
		return helper.Response(c, 500, "Gagal hapus alumni", nil)
	}
	return helper.Response(c, 200, "Alumni dihapus", nil)
}

// ====== LOGIKA BISNIS ======

// func CreateAlumnis(a models.Alumni) error {
// 	if a.NIM == "" || a.Nama == "" {
// 		return ErrInvalidData
// 	}
// 	return repository.InsertAlumni(a)
// }

// func UpdateAlumnis(id string, a models.Alumni) error {
// 	if a.NIM == "" || a.Nama == "" {
// 		return ErrInvalidData
// 	}
// 	return repository.UpdateAlumni(id, a)
// }

// func DeleteAlumnis(id string) error {
// 	return repository.DeleteAlumni(id)
// }
