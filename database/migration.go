package database

import (
	"crud-alumni/helper"
	"database/sql"
	"log"
)

// MigrateTesting hapus semua data dan insert ulang data sample
func MigrateTesting(DB *sql.DB) {
	// Hapus semua data (urutannya penting karena ada foreign key)
	_, err := DB.Exec(`
		DELETE FROM users;
	`)
	if err != nil {
		log.Fatalf("Gagal hapus data: %v", err)
	}
	adminPass := "123456"
	userPass := "123456"

	// Hash password
	adminHash, _ := helper.HashPassword(adminPass)
	userHash, _ := helper.HashPassword(userPass)

	// Insert admin
	_, err = DB.Exec(`
		INSERT INTO users (username, email, password_hash, role)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (username) DO NOTHING;
	`, "admin", "admin@university.com", adminHash, "admin")
	if err != nil {
		log.Fatalf("Gagal insert admin: %v", err)
	}

	// Insert user biasa
	_, err = DB.Exec(`
		INSERT INTO users (username, email, password_hash, role)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (username) DO NOTHING;
	`, "user1", "user1@university.com", userHash, "user")
	if err != nil {
		log.Fatalf("Gagal insert user1: %v", err)
	}

}
