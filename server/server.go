package main

import (
	"server/parserInterpreter/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	// Initialize default config, accept all origins
	app.Use(cors.New())

	// create group
	parserGroup := app.Group("/parser")

	// create route
	parserGroup.Post("/analyze", routes.AnalyzeAndParseCode())

	app.Listen(":3000")
}
