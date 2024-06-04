package bookcontroller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/herulobarto/go-restapi-fiber/models"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {

	var books []models.Book
	models.DB.Find(&books)

	return c.JSON(books)
}

func Show(c *fiber.Ctx) error {

	id := c.Params("id")
	var book models.Book
	if err := models.DB.First(&book, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Data tidak ditemukan",
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "data tidak ditemukan",
		})
	}

	return c.JSON(book)

}

func Create(c *fiber.Ctx) error {

	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := models.DB.Create(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(book)

}

func Update(c *fiber.Ctx) error {

	id := c.Params("id")

	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if models.DB.Where("id = ?", id).Updates(&book).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "tidak dapat megupdate data",
		})
	}

	return c.JSON(fiber.Map{
		"message": "data berhasil diupdate",
	})

}

func Delete(c *fiber.Ctx) error {
	return nil

}
