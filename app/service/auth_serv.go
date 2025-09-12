package service

import (
	"crud-alumni/helper"
	"crud-alumni/app/models"
	"crud-alumni/app/repository"
	"database/sql"
	"errors"
)

func Login(req models.LoginRequest) (*models.LoginResponse, error) {
	user, passwordHash, err := repository.FindUserByUsernameOrEmail(req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("username atau password salah")
		}
		return nil, errors.New("error database")
	}

	// cek password
	if !helper.CheckPassword(req.Password, passwordHash) {
		return nil, errors.New("username atau password salahh")
	}

	// generate token
	token, err := helper.GenerateToken(*user)
	if err != nil {
		return nil, errors.New("gagal generate token")
	}

	resp := &models.LoginResponse{
		User:  *user,
		Token: token,
	}

	return resp, nil
}
