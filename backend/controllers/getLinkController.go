package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kazion500/secure-share/models"
	"github.com/kazion500/secure-share/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetLink(c *fiber.Ctx) error {
	linkId := c.Params("linkId")

	if linkId == "" {
		return c.Status(400).JSON(
			types.ResponseType[types.ErrorType]{
				Success: false,
				Data:    types.ErrorType{Error: "id is required"},
			},
		)
	}

	linkIdObj, err := primitive.ObjectIDFromHex(linkId)

	if err != nil {
		return c.Status(400).JSON(
			types.ResponseType[types.ErrorType]{
				Success: false,
				Data:    types.ErrorType{Error: "id is invalid"},
			},
		)
	}

	link, err := models.FindLinkByShortLink(linkIdObj)

	if err != nil {
		return c.Status(500).JSON(
			types.ResponseType[types.ErrorType]{
				Success: false,
				Data:    types.ErrorType{Error: err.Error()},
			},
		)
	}

	if link.IsUsed {
		return c.Status(400).JSON(
			types.ResponseType[types.ErrorType]{
				Success: false,
				Data:    types.ErrorType{Error: "link has already been used"},
			},
		)
	}

	models.UpdateLink(linkIdObj, models.Link{
		IsUsed: true,
	})

	return c.JSON(
		types.ResponseType[models.Link]{
			Success: true,
			Data:    link,
		},
	)

}
