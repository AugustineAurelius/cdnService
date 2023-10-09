package handlers

import (
	"cdnService/internal/model"
	"encoding/base64"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func GetHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	logrus.Infof("Id %s", id)
	content, err := base64.StdEncoding.DecodeString(model.SN.Load(id))
	if err != nil {
		logrus.Fatalln("Failed to decode image data")
	}
	c.Set("Content-Type", "image/png")

	return c.Send([]byte(content))
}
