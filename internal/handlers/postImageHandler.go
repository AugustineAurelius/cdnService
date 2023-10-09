package handlers

import (
	"cdnService/internal/model"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"net/http"
)

func PostHandler(c *fiber.Ctx) error {
	content := new(model.Image)

	if err := c.BodyParser(&content); err != nil {
		logrus.Errorf("Error while body parse in POST handler")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	model.SN.Store(
		model.SN.Counter+1,
		content.ImageBase64,
	)
	return c.Status(http.StatusCreated).Send([]byte(
		fmt.Sprintf("Image uploaded successfully. URI ID: /image/%s", model.SN.GetCounter())))
}
