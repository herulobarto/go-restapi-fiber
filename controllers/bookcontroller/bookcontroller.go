package bookcontroller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/herulobarto/go-restapi-fiber/models"
)

func Index(c *fiber.Ctx) error {

	var books []models.Book
	models.DB.Find(&books)

	return c.Status(fiber.StatusOK).JSON(books)
}

func Show(c *fiber.Ctx) error {
	return nil

}

func Create(c *fiber.Ctx) error {
	return nil

}

func Update(c *fiber.Ctx) error {
	return nil

}

func Delete(c *fiber.Ctx) error {
	return nil

}
