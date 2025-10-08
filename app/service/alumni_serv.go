package service

import (
	"crud-alumni/app/models"
	"crud-alumni/app/repository"
	"crud-alumni/helper"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// ====== HANDLER ======

// func GetAlumniByFakultasHandler(c *fiber.Ctx) error {
// 	fak := c.Params("fakul")

// 	data, err := GetAllAlumniByFak(fak)
// 	if err != nil {
// 		return helper.Response(c, 400, err.Error(), nil)
// 	}
// 	return helper.Response(c, 200, "OK", data)
// }

func GetAlumniListHandler(c *fiber.Ctx) error {
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
}

func GetAlumniByIDHandler(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	data, err := GetAlumniByID(id)
	if err != nil {
		return helper.Response(c, 404, "Alumni tidak ditemukan", nil)
	}
	return helper.Response(c, 200, "OK", data)
}

func CreateAlumniHandler(c *fiber.Ctx) error {
	var a models.Alumni
	if err := c.BodyParser(&a); err != nil {
		return helper.Response(c, 400, "Input tidak valid", nil)
	}
	if err := CreateAlumni(a); err != nil {
		return helper.Response(c, 400, err.Error(), nil)
	}
	return helper.Response(c, 201, "Alumni ditambahkan", a)
}

func UpdateAlumniHandler(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var a models.Alumni
	if err := c.BodyParser(&a); err != nil {
		return helper.Response(c, 400, "Input tidak valid", nil)
	}
	if err := UpdateAlumni(id, a); err != nil {
		return helper.Response(c, 400, err.Error(), nil)
	}
	return helper.Response(c, 200, "Alumni diupdate", a)
}

func DeleteAlumniHandler(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := DeleteAlumni(id); err != nil {
		return helper.Response(c, 500, "Gagal hapus alumni", nil)
	}
	return helper.Response(c, 200, "Alumni dihapus", nil)
}

// ====== LOGIKA BISNIS ======

// func GetAllAlumniByFak(fak string) ([]models.Alumni, error) {
// 	if fak == "" {
// 		return nil, ErrInvalidData
// 	}
// 	return repository.GetAllAlumniByFak(fak)
// }

func GetAlumniByID(id int) (models.Alumni, error) {
	return repository.GetAlumniByID(id)
}

func CreateAlumni(a models.Alumni) error {
	if a.NIM == "" || a.Nama == "" || a.Email == "" {
		return ErrInvalidData
	}
	return repository.InsertAlumni(a)
}

func UpdateAlumni(id int, a models.Alumni) error {
	if a.NIM == "" || a.Nama == "" || a.Email == "" {
		return ErrInvalidData
	}
	return repository.UpdateAlumni(id, a)
}

func DeleteAlumni(id int) error {
	return repository.DeleteAlumni(id)
}
