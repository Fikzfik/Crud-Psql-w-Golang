package route

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	// daftar route dipisah ke file lain
	RegisterAlumniRoutes(api)
	RegisterPekerjaanRoutes(api)
	RegisterAuthRoutes(api)
}
