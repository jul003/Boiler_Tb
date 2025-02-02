package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/jul003/Boiler_Tb/config"
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jul003/Boiler_Tb/url"
	_ "github.com/jul003/Boiler_Tb/docs"
)

// @title TES SWAGGER ULBI
// @version 1.0
// @description This is a sample swagger for Fiber

// @contact.name API Support
// @contact.url https://github.com/jul003
// @contact.email indra@ulbi.ac.id

// @host tb-2024-2df3fc45e2ad.herokuapp.com
// @BasePath /
// @schemes https http

func main() {
	// Initialize Fiber app with config
	site := fiber.New(config.Iteung)

	// Enable CORS middleware using config
	site.Use(cors.New(config.Cors))

	// Define routes using Fiber's own methods
	url.Web(site)

	// Start the server
	log.Fatal(site.Listen(musik.Dangdut()))

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())
}
