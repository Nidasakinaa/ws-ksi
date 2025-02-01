package url

import (
	"github.com/Nidasakinaa/ws-kaloriku/controller"
	"github.com/Nidasakinaa/ws-kaloriku/handler"
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

	page.Get("/checkip", controller.Homepage)
	page.Get("/menu", controller.GetMenu)
	page.Get("/menu/:id", controller.GetMenuID)
	page.Get("/menu/:category", controller.GetMenuByCategory)
	page.Post("/insert", controller.InsertDataMenu)
	page.Put("/update/:id", controller.UpdateDataMenuItem)
	page.Delete("/delete/:id", controller.DeleteMenuItemByID)

	page.Get("/user", controller.GetUser)
	page.Get("/user/:id", controller.GetUserID)
	page.Post("/insert", controller.InsertDataUser)
	page.Put("/user/update/:id", controller.UpdateDataUser)
	page.Delete("/user/delete/:id", controller.DeleteUserByID)
	page.Post("/registeruser", handler.Register)
	page.Post("/login", handler.Login)
	page.Post("/loginCust", handler.CustomerLogin)

	page.Get("/docs/*", swagger.HandlerDefault)

	// page.Use(middleware.AuthMiddleware())
	page.Get("/dashboard", handler.DashboardPage)
}
