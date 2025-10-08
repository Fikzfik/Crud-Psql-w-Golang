package service

import (
	"crud-alumni/app/models"
	"crud-alumni/app/repository"
	"crud-alumni/helper"
	"database/sql"
	"errors"

	"github.com/gofiber/fiber/v2"
)

// ===== HANDLERS =====

func LoginHandler(c *fiber.Ctx) error {
	var req models.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return helper.Response(c, 400, "Request body tidak valid", nil)
	}

	resp, err := Login(req)
	if err != nil {
		return helper.Response(c, 401, err.Error(), nil)
	}
	return helper.Response(c, 200, "Login berhasil", resp)
}

func ProfileHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(int)
	username := c.Locals("username").(string)
	role := c.Locals("role").(string)

	return helper.Response(c, 200, "Profile berhasil diambil", fiber.Map{
		"user_id":  userID,
		"username": username,
		"role":     role,
	})
}

func IsDeletedProfileHandler(c *fiber.Ctx) error {
	role := c.Locals("role").(string)
	userID := c.Locals("user_id").(int)

	err := IsDeleted(role, userID)
	if err != nil {
		return helper.Response(c, 500, "Gagal hapus", nil)
	}
	return helper.Response(c, 200, "Akun dihapus", nil)
}

// ===== LOGIKA BISNIS =====

func Login(req models.LoginRequest) (*models.LoginResponse, error) {
	user, passwordHash, err := repository.FindUserByUsernameOrEmail(req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("username atau password salah")
		}
		return nil, errors.New("error database")
	}

	if !helper.CheckPassword(req.Password, passwordHash) {
		return nil, errors.New("username atau password salah")
	}

	token, err := helper.GenerateToken(*user)
	if err != nil {
		return nil, errors.New("gagal generate token")
	}

	return &models.LoginResponse{
		User:  *user,
		Token: token,
	}, nil
}
