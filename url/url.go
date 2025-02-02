package url

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/swagger"
	"github.com/jul003/Boiler_Tb/controller"
)

func Web(page *fiber.App) {
	// Initialize session store
	store := session.New()

	// Use session middleware
	page.Use(func(c *fiber.Ctx) error {
		// Store session in context
		c.Locals("session", store)
		return c.Next()
	})

	// CSRF protection middleware
	page.Use(csrf.New(csrf.Config{
		KeyLookup:      "cookie:csrf_token",
		CookieName:     "csrf_token",
		CookieSecure:   false, // Set `true` jika pakai HTTPS
		CookieHTTPOnly: false, // Jika frontend perlu membaca cookie langsung
	}))
	

	// Endpoint to get CSRF token
	page.Get("/csrf-token", func(c *fiber.Ctx) error {
		csrfToken := c.Cookies("csrf_token") // Ambil dari cookie
		if csrfToken == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "CSRF token not found"})
		}
		return c.JSON(fiber.Map{"csrf_token": csrfToken})
	})
	

	// Other routes
	page.Post("/", controller.Sink)
	page.Put("/", controller.Sink)
	page.Patch("/", controller.Sink)
	page.Delete("/", controller.Sink)

	page.Get("/gadget", controller.GetAllGagdet)
	page.Get("/review", controller.GetAllReview)

	// Routes for inserting gadgets and reviews with CSRF protection
	page.Post("/insertgadget", controller.InsertDataGadget)
	page.Post("/insertreview", controller.InsertDataReview)

	// Routes for updating and deleting gadgets
	page.Put("/updategadget/:id", controller.UpdateDataGadget)
	page.Delete("/deletegadget/:id", controller.DeleteGadgetByID)

	// Swagger documentation route
	page.Get("/docs/*", swagger.HandlerDefault)

	// Route to get gadget by ID
	page.Get("/gadget/:id", controller.GetGadgetByID)
}
