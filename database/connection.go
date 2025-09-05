package database

import (
	"database/sql"
	"fmt"
	"log"

	"crud-alumni/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	host := config.GetEnv("DB_HOST", "localhost")
	port := config.GetEnv("DB_PORT", "5432")
	user := config.GetEnv("DB_USER", "postgres")
	pass := config.GetEnv("DB_PASS", "postgres")
	name := config.GetEnv("DB_NAME", "alumni_db")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, name)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(" Gagal koneksi DB:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(" DB tidak bisa di-ping:", err)
	}

	fmt.Println(" Database connected")
}
