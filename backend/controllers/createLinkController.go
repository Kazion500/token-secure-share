package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kazion500/secure-share/models"
	"github.com/kazion500/secure-share/types"
	"github.com/kazion500/secure-share/utils"
)

func CreateLink(c *fiber.Ctx) error {
	var body types.RequestBody

	if err := c.BodyParser(&body); err != nil {
		return c.Status(422).JSON(
			types.ResponseType[types.ErrorType]{
				Success: false,
				Data:    types.ErrorType{Error: "data is required"},
			},
		)
	}

	if body.Data == "" {
		return c.Status(400).JSON(
			types.ResponseType[types.ErrorType]{
				Success: false,
				Data:    types.ErrorType{Error: "data is required"},
			},
		)
	}

	shortCode := utils.GenerateShortCode(6)

	newLink := models.Link{
		IsViewed:    false,
		ShortCode:   shortCode,
		LinkContent: body.Data,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	link, err := models.InsertLink(newLink)

	if err != nil {
		return c.Status(500).JSON(
			types.ResponseType[types.ErrorType]{
				Success: false,
				Data:    types.ErrorType{Error: "internal server error"},
			},
		)
	}

	return c.JSON(
		types.ResponseType[types.DataType]{
			Success: true,
			Data: types.DataType{
				Link: link.ShortCode,
			},
		},
	)
}
