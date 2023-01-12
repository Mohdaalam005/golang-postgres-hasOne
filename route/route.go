package route

import (
	"github.com/gofiber/fiber"
	"github.com/mohdaalam005/one-to-one/controller"
)

func Route(app *fiber.App) {
	app.Get("/users", controller.GetAllUser)
	app.Get("/users/:id", controller.GetUser)
	app.Delete("/users/:id", controller.DeleteUser)
	app.Put("/users/:id", controller.UpdateUsers)
	app.Post("/users", controller.CreateUser)

	app.Listen(":8080")
}
