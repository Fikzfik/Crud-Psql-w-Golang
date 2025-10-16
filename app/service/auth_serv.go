package service

import (
	"crud-alumni/app/models"
	"crud-alumni/app/repository"
	"crud-alumni/helper"
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
	userID := c.Locals("user_id").(string)
	email := c.Locals("email").(string)
	role := c.Locals("role").(string)

	return helper.Response(c, 200, "Profile berhasil diambil", fiber.Map{
		"user_id": userID,
		"email":   email,
		"role":    role,
	})
}

// ===== LOGIKA BISNIS =====

func Login(req models.LoginRequest) (*models.LoginResponse, error) {
	user, err := repository.FindUserByEmail(req.Email)
	if err != nil {
		return nil, errors.New("email tidak ditemukan atau salah")
	}

	if !helper.CheckPassword(req.Password, user.PasswordHash) {
		return nil, errors.New("password salah")
	}

	// Convert ObjectID ke string
	token, err := helper.GenerateToken(models.User{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	})
	if err != nil {
		return nil, errors.New("gagal generate token")
	}

	return &models.LoginResponse{
		User: models.User{
			ID:        user.ID,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
		},
		Token: token,
	}, nil
}
