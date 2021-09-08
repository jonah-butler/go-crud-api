package router

import (
	"github.com/gofiber/fiber/v2"
)

// func Router(str string) {
// 	fmt.Println(str)
// }

func HelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World")
}

func LandingRoutes(app *fiber.App) {
	app.Get("/", HelloWorld)
}
