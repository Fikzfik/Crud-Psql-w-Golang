package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"crud-alumni/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database // Ganti dari *sql.DB ke *mongo.Database

// ConnectDB membuat koneksi ke MongoDB
func ConnectDB() {
	// Ambil konfigurasi dari .env
	mongoURI := config.GetEnv("MONGO_URI", "mongodb://localhost:27017")
	dbName := config.GetEnv("MONGO_DB_NAME", "alumni_db")

	// Siapkan client MongoDB
	clientOptions := options.Client().ApplyURI(mongoURI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("❌ Gagal koneksi ke MongoDB: %v", err)
	}

	// Ping MongoDB untuk memastikan koneksi
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("❌ MongoDB tidak bisa di-ping: %v", err)
	}

	fmt.Println("✅ Berhasil terhubung ke MongoDB!")

	// Simpan database global
	DB = client.Database(dbName)
}
