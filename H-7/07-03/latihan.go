package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Contact struct {
	Name     string
	Location string
	Email    string
}

var contacts []Contact

func homeHandler(c *fiber.Ctx) error {
	return c.SendString("Welcome to the Home Page!")
}

func contactHandler(c *fiber.Ctx) error {
	result := "Contacts:\n"
	for _, contact := range contacts {
		result += fmt.Sprintf("Name: %s, Location: %s, Email: %s\n", contact.Name, contact.Location, contact.Email)
	}
	return c.SendString(result)
}

func addContactHandler(c *fiber.Ctx) error {
	if c.Method() == fiber.MethodPost {
		name := c.FormValue("name")
		location := c.FormValue("location")
		email := c.FormValue("email")

		newContact := Contact{Name: name, Location: location, Email: email}
		contacts = append(contacts, newContact)
		return c.SendString("Contact added successfully!")
	}

	return fiber.ErrMethodNotAllowed
}

func main() {
	app := fiber.New()

	app.Get("/", homeHandler)
	app.Get("/contact", contactHandler)
	app.Post("/addcontact", addContactHandler)

	fmt.Println("Server is running on http://localhost:8080")
	app.Listen(":8080")
}
