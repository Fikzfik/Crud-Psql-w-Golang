package main

import (
	"crud-alumni/config"
	"crud-alumni/database"
	"crud-alumni/route"
)

func main() {
	config.LoadEnv()
	config.InitLogger()
	database.ConnectDB()

	app := config.NewApp()
	route.RegisterRoutes(app)

	port := config.GetEnv("APP_PORT", "3000")
	app.Listen(":" + port)
}