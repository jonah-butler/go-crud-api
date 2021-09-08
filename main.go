package main

import (
	db "go-vue-auth/db"
	routes "go-vue-auth/router"
	userRoutes "go-vue-auth/router/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"os"
	//"github.com/joho/godotenv"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func initDatabase() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + name + " port=" + port + " sslmode=disable"
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
