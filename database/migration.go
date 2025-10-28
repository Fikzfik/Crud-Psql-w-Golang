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
	pekerjaanCol := DB.Collection("pekerjaan_alumni")

	// 🔹 Bersihkan koleksi lama
	userCol.DeleteMany(ctx, bson.M{})
	alumniCol.DeleteMany(ctx, bson.M{})
	pekerjaanCol.DeleteMany(ctx, bson.M{})

	// 🔹 Hash password
	adminHash, _ := helper.HashPassword("123456")
	userHash, _ := helper.HashPassword("123456")

	// 🔹 Data users
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

	_, err := userCol.InsertMany(ctx, []interface{}{adminUser, user1, user2})
	if err != nil {
		log.Fatalf("❌ Gagal insert users: %v", err)
	}

	// 🔹 Data alumni
	alumni1 := models.Alumni{
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
	}
	alumni2 := models.Alumni{
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
	}
	alumni3 := models.Alumni{
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
	}

	_, err = alumniCol.InsertMany(ctx, []interface{}{alumni1, alumni2, alumni3})
	if err != nil {
		log.Fatalf("❌ Gagal insert alumni: %v", err)
	}

	// 🔹 Data pekerjaan alumni
	pekerjaanData := []interface{}{
		models.PekerjaanAlumni{
			ID:                 primitive.NewObjectID(),
			AlumniID:           alumni1.ID,
			NamaPerusahaan:     "Universitas ABC",
			PosisiJabatan:      "Administrator Sistem",
			BidangIndustri:     "Pendidikan",
			LokasiKerja:        "Jakarta",
			GajiRange:          "Rp7.000.000 - Rp9.000.000",
			TanggalMulaiKerja:  time.Date(2023, 8, 1, 0, 0, 0, 0, time.Local),
			StatusPekerjaan:    "Tetap",
			DeskripsiPekerjaan: "Mengelola server dan infrastruktur TI universitas.",
			IsDeleted:          false,
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		},
		models.PekerjaanAlumni{
			ID:                 primitive.NewObjectID(),
			AlumniID:           alumni2.ID,
			NamaPerusahaan:     "PT Teknologi Nusantara",
			PosisiJabatan:      "Data Analyst",
			BidangIndustri:     "Teknologi Informasi",
			LokasiKerja:        "Bandung",
			GajiRange:          "Rp8.000.000 - Rp12.000.000",
			TanggalMulaiKerja:  time.Date(2024, 7, 15, 0, 0, 0, 0, time.Local),
			StatusPekerjaan:    "Kontrak",
			DeskripsiPekerjaan: "Menganalisis data pelanggan untuk mendukung keputusan bisnis.",
			IsDeleted:          false,
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		},
		models.PekerjaanAlumni{
			ID:                 primitive.NewObjectID(),
			AlumniID:           alumni3.ID,
			NamaPerusahaan:     "CV Mekanika Cerdas",
			PosisiJabatan:      "Teknisi Komputer",
			BidangIndustri:     "Elektronik",
			LokasiKerja:        "Surabaya",
			GajiRange:          "Rp5.000.000 - Rp7.000.000",
			TanggalMulaiKerja:  time.Date(2025, 2, 1, 0, 0, 0, 0, time.Local),
			StatusPekerjaan:    "Magang",
			DeskripsiPekerjaan: "Melakukan perawatan dan perbaikan perangkat komputer.",
			IsDeleted:          false,
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		},
	}

	_, err = pekerjaanCol.InsertMany(ctx, pekerjaanData)
	if err != nil {
		log.Fatalf("❌ Gagal insert pekerjaan_alumni: %v", err)
	}

	fmt.Println("✅ Sample data berhasil dimigrasikan ke MongoDB:")
	fmt.Println("   → 3 users")
	fmt.Println("   → 3 alumni")
	fmt.Println("   → 3 pekerjaan_alumni")
}
