package url

import (
	"github.com/jul003/Boiler_Tb/controller"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/swagger"
)

func Web(page *fiber.App) {
	// page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	// page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth

	page.Get("/", controller.Sink)
	page.Post("/", controller.Sink)
	page.Put("/", controller.Sink)
	page.Patch("/", controller.Sink)
	page.Delete("/", controller.Sink)
	page.Options("/", controller.Sink)


	page.Get("/gadget", controller.GetAllGagdet)
	page.Get("/review", controller.GetAllReview) 
	page.Post("/insertgadget", controller.InsertDataGadget)
	page.Post("/insertreview", controller.InsertDataReview)
	page.Put("/updategadget/:id", controller.UpdateDataGadget)
	page.Delete("/deletegadget/:id", controller.DeleteGadgetByID)
	page.Get("/docs/*", swagger.HandlerDefault)

	page.Get("/gadget/:id", controller.GetGadgetByID)

	
}
