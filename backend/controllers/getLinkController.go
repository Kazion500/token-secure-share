package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kazion500/secure-share/models"
	"github.com/kazion500/secure-share/types"
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

	link, err := models.FindLinkByShortCode(linkId)

	if err != nil {
		return c.Status(500).JSON(
			types.ResponseType[types.ErrorType]{
				Success: false,
				Data:    types.ErrorType{Error: err.Error()},
			},
		)
	}

	if link.IsViewed {
		return c.Status(400).JSON(
			types.ResponseType[types.ErrorType]{
				Success: false,
				Data:    types.ErrorType{Error: "link has already been used"},
			},
		)
	}

	models.UpdateLink(link.ID, models.Link{
		IsViewed: true,
	})

	return c.JSON(
		types.ResponseType[types.DataType]{
			Success: true,
			Data: types.DataType{
				Content: link.LinkContent,
			},
		},
	)

}
