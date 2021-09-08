package user

import (
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	app.Get("/users", getAllUsers)
	app.Get("/user/:id", getUser)
	app.Post("/user", newUser)
	app.Delete("/user", deleteUser)
	app.Put("/user", updateUser)
}
