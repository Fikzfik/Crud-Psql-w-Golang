package database

import (
	"crud-alumni/helper"
	"database/sql"
	"log"
)

// MigrateTesting hapus semua data dan insert ulang data sample
func MigrateTesting(DB *sql.DB) {
	// Hapus semua data (urutannya penting karena ada foreign key)
	
	adminPass := "123456"
	userPass := "123456"

	// Hash password
	adminHash, _ := helper.HashPassword(adminPass)
	userHash, _ := helper.HashPassword(userPass)

	// Insert admin
	_, err := DB.Exec(`
		INSERT INTO users (username,alumni_id, email, password_hash, role)
		VALUES ($1, $2, $3, $4,$5)
		ON CONFLICT (username) DO NOTHING;
	`, "admin",1, "admin1@university.com", adminHash, "admin")
	if err != nil {
		log.Fatalf("Gagal insert admin: %v", err)
	}

	// Insert user biasa
	_, err = DB.Exec(`
		INSERT INTO users (username, alumni_id,email, password_hash, role)
		VALUES ($1, $2, $3, $4,$5)
		ON CONFLICT (username) DO NOTHING;
	`, "user4",2, "user3@university.com", userHash, "user")
	if err != nil {
		log.Fatalf("Gagal insert user1: %v", err)
	}

	_, err = DB.Exec(`
		INSERT INTO users (username, alumni_id,email, password_hash, role)
		VALUES ($1, $2, $3, $4,$5)
		ON CONFLICT (username) DO NOTHING;
	`, "user4",3, "user4@university.com", userHash, "user")
	if err != nil {
		log.Fatalf("Gagal insert user1: %v", err)
	}
}
