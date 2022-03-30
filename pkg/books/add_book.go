package books

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hellokvn/go-fiber-api/pkg/common/models"
)

type AddBookRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (h handler) AddBook(c *fiber.Ctx) error {
	body := AddBookRequestBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var book models.Book

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	if result := h.DB.Create(&book); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.JSON(&book)
}
