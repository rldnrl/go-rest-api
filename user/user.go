package user

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var error error

const DNS = "root:admin@tcp(127.0.0.1:3306)/godb?charset=utf8"

type User struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func InitialMigration() {
	DB, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to Database")
	}

	DB.AutoMigrate(&User{})
}

func GetUsers(c *fiber.Ctx) error {
	var users []User
	DB.Find(&users)
	return c.JSON(&users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User
	DB.Find(&user, id)
	return c.JSON(&user)
}

func SaveUser(c *fiber.Ctx) error {
	user := new(User)
	err := c.BodyParser(user)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	DB.Save(&user)
	return c.JSON(&user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User
	DB.First(&user, id)

	if user.Email == "" {
		return c.Status(400).SendString("User not available")
	}

	DB.Delete(&user)
	return c.SendString("User is deleted")
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(User)
	DB.First(&user, id)

	if user.Email == "" {
		return c.Status(400).SendString("User not available")
	}

	err := c.BodyParser(&user)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	DB.Save(&user)

	return c.JSON(&user)
}
