package main

import (
	"github.com/gofiber/fiber/v2"
)

type User struct {
	FirstName string
	LastName  string
}

func handleUser(c *fiber.Ctx) error {
	user := User{
		FirstName: "Daniel",
		LastName:  "Reyes",
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func handleCreateUser(c *fiber.Ctx) error {

	user := User{}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(user)

}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/user", handleUser)
	app.Post("/user", handleCreateUser)

	app.Listen(":3000")
}
