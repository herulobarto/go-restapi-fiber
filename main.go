package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/herulobarto/go-restapi-fiber/controllers/bookcontroller"
	"github.com/herulobarto/go-restapi-fiber/models"
)

func main() {
	models.ConnectDatabase()

	app := fiber.New()

	// Grouping routes /api/books/...
	api := app.Group("/api")
	book := api.Group("/books")

	book.Get("/", bookcontroller.Index)
	book.Get("/:id", bookcontroller.Show)
	book.Post("/", bookcontroller.Create)
	book.Put("/:id", bookcontroller.Update)
	book.Delete("/:id", bookcontroller.Delete)

	app.Listen(":8000")
}
