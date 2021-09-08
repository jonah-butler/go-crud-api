package user

import (
	db "go-vue-auth/db"

	"github.com/gofiber/fiber/v2"

	f "fmt"
)

type User struct {
	// gorm.Model
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func getAllUsers(c *fiber.Ctx) error {
	db := db.DBConn
	var users []User
	db.Find(&users)
	f.Println(users)
	f.Println(users[0].Email)
	return c.SendString("Hello from all Users")
}

func getUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	db := db.DBConn
	var user User
	if result := db.Where("id = ?", userId).First(&user); result.Error != nil {
		f.Println(result.Error)
	} else {
		f.Println("this", user)
	}
	return c.SendString("Hello from a single user: " + c.Params("id"))
}

func newUser(c *fiber.Ctx) error {
	user := &User{}
	err := c.BodyParser(user)
	if err != nil {
		f.Printf("body parser error: %v", err)
		return c.SendString("error")
	}
	db := db.DBConn
	db.Select("Email", "Password").Create(&user)
	return c.JSON(user)
}

func updateUser(c *fiber.Ctx) error {
	return c.SendString("Hello from update user")
}

func deleteUser(c *fiber.Ctx) error {
	return c.SendString("Hello from delete user")
}
