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

func GetAlumniListHandler(c *fiber.Ctx) error {
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

func GetAlumniByIDHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	data, err := repository.GetAlumniByID(id)
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
	id := c.Params("id")
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
	id := c.Params("id")
	if err := DeleteAlumni(id); err != nil {
		return helper.Response(c, 500, "Gagal hapus alumni", nil)
	}
	return helper.Response(c, 200, "Alumni dihapus", nil)
}

// ====== LOGIKA BISNIS ======

func CreateAlumni(a models.Alumni) error {
	if a.NIM == "" || a.Nama == "" {
		return ErrInvalidData
	}
	return repository.InsertAlumni(a)
}

func UpdateAlumni(id string, a models.Alumni) error {
	if a.NIM == "" || a.Nama == "" {
		return ErrInvalidData
	}
	return repository.UpdateAlumni(id, a)
}

func DeleteAlumni(id string) error {
	return repository.DeleteAlumni(id)
}
