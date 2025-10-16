package database

import (
	"context"
	"crud-alumni/app/models"
	"crud-alumni/helper"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MigrateTesting akan menghapus semua data lama dan menambahkan data contoh baru ke MongoDB
func MigrateTesting(DB *mongo.Database) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userCol := DB.Collection("users")
	alumniCol := DB.Collection("alumni")

	// 🔹 Hapus data lama (biar bersih)
	userCol.DeleteMany(ctx, bson.M{})
	alumniCol.DeleteMany(ctx, bson.M{})

	// 🔹 Hash password
	adminHash, _ := helper.HashPassword("123456")
	userHash, _ := helper.HashPassword("123456")

	// 🔹 Buat data contoh users
	adminUser := models.User{
		ID:           primitive.NewObjectID(),
		Email:        "admin1@university.com",
		PasswordHash: adminHash,
		Role:         "admin",
		CreatedAt:    time.Now(),
	}

	user1 := models.User{
		ID:           primitive.NewObjectID(),
		Email:        "user1@university.com",
		PasswordHash: userHash,
		Role:         "user",
		CreatedAt:    time.Now(),
	}

	user2 := models.User{
		ID:           primitive.NewObjectID(),
		Email:        "user2@university.com",
		PasswordHash: userHash,
		Role:         "user",
		CreatedAt:    time.Now(),
	}

	// 🔹 Masukkan data users
	_, err := userCol.InsertMany(ctx, []interface{}{adminUser, user1, user2})
	if err != nil {
		log.Fatalf("❌ Gagal insert users: %v", err)
	}

	// 🔹 Buat data alumni (relasi ke user_id)
	alumniData := []interface{}{
		models.Alumni{
			ID:         primitive.NewObjectID(),
			UserID:     adminUser.ID,
			NIM:        "A001",
			Nama:       "Admin Utama",
			Jurusan:    "Teknik Informatika",
			Fakultas:   "FTI",
			Angkatan:   2019,
			TahunLulus: 2023,
			NoTelepon:  "081234567890",
			Alamat:     "Jl. Kampus No.1",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		models.Alumni{
			ID:         primitive.NewObjectID(),
			UserID:     user1.ID,
			NIM:        "B002",
			Nama:       "Fikri Ardi",
			Jurusan:    "Sistem Informasi",
			Fakultas:   "FTI",
			Angkatan:   2020,
			TahunLulus: 2024,
			NoTelepon:  "081312345678",
			Alamat:     "Jl. Merpati No.9",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		models.Alumni{
			ID:         primitive.NewObjectID(),
			UserID:     user2.ID,
			NIM:        "C003",
			Nama:       "Budi Santoso",
			Jurusan:    "Teknik Komputer",
			Fakultas:   "FTI",
			Angkatan:   2021,
			TahunLulus: 2025,
			NoTelepon:  "081390000123",
			Alamat:     "Jl. Anggrek No.3",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	_, err = alumniCol.InsertMany(ctx, alumniData)
	if err != nil {
		log.Fatalf("❌ Gagal insert alumni: %v", err)
	}

	fmt.Println("✅ Sample data berhasil dimigrasikan ke MongoDB:")
	fmt.Println("   → 3 users")
	fmt.Println("   → 3 alumni")
}
