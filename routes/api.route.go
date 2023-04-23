package routes

import (
	"jobhun-intern/controllers"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(server *fiber.App) {
	server.Get("/", controllers.SendHello)
}
