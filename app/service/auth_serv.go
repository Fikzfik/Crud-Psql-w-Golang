package service

import (
	"crud-alumni/app/models"
	"crud-alumni/app/repository"
	"crud-alumni/helper"
	"errors"

	"github.com/gofiber/fiber/v2"
)

// ===== HANDLERS =====

// @Summary Login user
// @Description Login dengan email dan password untuk mendapatkan token JWT
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body models.LoginRequest true "Data login (email & password)"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} map[string]interface{} "Request body tidak valid"
// @Failure 401 {object} map[string]interface{} "Email atau password salah"
// @Router /login [post]
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

// @Summary Mendapatkan profil user yang sedang login
// @Description Mengambil data profil berdasarkan token JWT yang dikirim di header Authorization
// @Tags Auth
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{} "Profile berhasil diambil"
// @Failure 401 {object} map[string]interface{} "Token tidak valid atau tidak ada"
// @Router /profile [get]
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
