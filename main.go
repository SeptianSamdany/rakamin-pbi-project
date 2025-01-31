package main

import (
	"log"
	"rakamin-project/config"
	"rakamin-project/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Inisialisasi database
	db := config.InitDB()
	if db == nil {
		log.Fatal("Failed to connect to database")
	}

	// Inisialisasi Fiber
	app := fiber.New()

	// Setup Routes
	routes.SetupRoutes(app, db)

	// Jalankan server
	log.Println("Server running on port 8080")
	log.Fatal(app.Listen(":8080"))
}
