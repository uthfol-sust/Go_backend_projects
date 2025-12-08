package main

import (
	"log"
	"os"
	"urlshorten/routes"

	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/url/:short", routes.ResolverURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	godotenv.Load()

	app := fiber.New()

	app.Use(fiberLogger.New())

	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
