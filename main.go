package main

import (
	db "go-vue-auth/db"
	routes "go-vue-auth/router"
	userRoutes "go-vue-auth/router/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	//"os"
	//"github.com/joho/godotenv"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func initDatabase() {
	var err error
	dsn := "host=localhost user=postgres password=Porkpie666 dbname=go-test port=5432 sslmode=disable"
	db.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("error connecting to database")
	} else {
		fmt.Printf("successfully connected to databse:")
	}
}

func main() {
	initDatabase()

	app := fiber.New()

	routes.LandingRoutes(app)
	userRoutes.SetupUserRoutes(app)

	// app.Get("/", func(c *fiber.Ctx) error {
	// user := Users{Id: 2, Email: "test@test.com", Password: "testyotest"}
	// // db.First(&user)
	// db.Select("Id", "Email", "Password").Create(&user)
	// fmt.Println("result is ", user)
	// 	fmt.Println(db)
	// 	return c.SendString("yo got it")
	// })
	app.Listen(":8000")
}
