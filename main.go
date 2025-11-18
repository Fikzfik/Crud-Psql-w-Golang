package main

import (
	"crud-alumni/config"
	"crud-alumni/database"
	_ "crud-alumni/docs" // wajib untuk swagger
	"crud-alumni/route"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

// @title CRUD Alumni API
// @version 1.0
// @description Masukkan token JWT dengan format: Bearer <token>
// @description API untuk manajemen data alumni, pekerjaan, dan user
// @host localhost:3000
// @BasePath /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	// === Load konfigurasi dasar ===
	config.LoadEnv()
	config.InitLogger()

	// === Koneksi ke Database ===
	database.ConnectDB()
	// database.MigrateTesting(database.DB) // uncomment jika perlu

	// === Inisialisasi Aplikasi Fiber ===
	app := config.NewApp()

	// === Middleware CORS ===
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000, http://127.0.0.1:3000",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: false, // ⚠️ HARUS false kalau pakai wildcard / domain sama
	}))

	// === Registrasi Route ===
	route.RegisterRoutes(app)

	// === Swagger Documentation ===
	app.Get("/swagger/*", swagger.HandlerDefault) // akses: localhost:3000/swagger/index.html

	// === Jalankan Server ===
	port := config.GetEnv("APP_PORT", "3000")
	app.Listen(":" + port)
}
