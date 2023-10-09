package main

import (
	"cdnService/internal/config"
	"cdnService/internal/handlers"
	"cdnService/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

var cfg config.Config
var router = fiber.New()

func init() {

	var err error
	cfg, err = config.LoadConfig()
	if err != nil {
		logrus.Fatal("Couldn't parse config", err)
	}
	model.SN = model.NewSyncMap()

	router.Post("/upload", handlers.PostHandler)
	router.Get("/image/:id", handlers.GetHandler)
}

func main() {

	model.SN.Store("1", "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")

	err := router.Listen(cfg.CDNPort) // инизиализируем сервер
	if err != nil {
		logrus.Fatal("Could not to start listen on port : ", cfg.CDNPort)
	}
}
