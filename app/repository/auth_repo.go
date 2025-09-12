package repository

import (
	"crud-alumni/database"
	"crud-alumni/app/models"
	"database/sql"
)

func FindUserByUsernameOrEmail(identifier string) (*models.User, string, error) {
	var user models.User
	var passwordHash string

	err := database.DB.QueryRow(`
		SELECT id, username, email, password_hash, role, created_at
		FROM users
		WHERE username = $1 OR email = $1
	`, identifier).Scan(
		&user.ID, &user.Username, &user.Email, &passwordHash, &user.Role, &user.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, "", sql.ErrNoRows
		}
		return nil, "", err
	}

	return &user, passwordHash, nil
}
