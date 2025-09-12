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
	database.MigrateTesting(database.DB)
	
	app := config.NewApp()
	route.RegisterRoutes(app)

	port := config.GetEnv("APP_PORT", "3000")
	app.Listen(":" + port)
}
