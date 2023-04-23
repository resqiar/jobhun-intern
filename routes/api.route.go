package routes

import (
	"jobhun-intern/controllers"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(server *fiber.App) {
	server.Get("/", controllers.SendHello)

	v1 := server.Group("v1")
	v1.Get("/mahasiswa/list", controllers.SendAllMahasiswa)
}
