package handler

import (
	"app/db"
	"app/middleware"
	"app/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func RouteProducts(server fiber.Router) {
	router := server.Group("/products")
	router.Get("", GetAllProducts)
	router.Post("", middleware.Auth, CreateProduct)
	router.Delete(":id", middleware.Auth, DeleteProductByID)
}

func GetAllProducts(c *fiber.Ctx) error {
	products := []model.Product{}
	err := db.Con.Debug().Order("id DESC").Find(&products).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.JSON(fiber.Map{
		"data": products,
	})
}

func CreateProduct(c *fiber.Ctx) error {
	user := c.Locals("email")
	fmt.Println("CreateProduct", user)

	in := &model.Product{}
	if err := c.BodyParser(&in); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := db.Con.Create(&in).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.JSON(fiber.Map{
		"message": "created",
	})
}

func DeleteProductByID(c *fiber.Ctx) error {
	id := c.Params("id")

	err := db.Con.Where("id = ?", id).Delete(&model.Product{}).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.JSON(fiber.Map{
		"message": "deleted",
	})
}
