package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rs/zerolog"
	"kino/internal/shared/config"
	"kino/internal/shared/log"
	"kino/internal/shared/repository"
)

type Handler struct {
	repository *repository.Repository
	logger     *zerolog.Logger
	conf       *config.Config
}

func NewHandler(repository *repository.Repository, logger *zerolog.Logger, conf *config.Config) *Handler {
	return &Handler{repository: repository, logger: logger, conf: conf}
}

func (h *Handler) InitRouter() {
	f := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
	})

	f.Use(cors.New(cors.Config{
		AllowOrigins: h.conf.Application.ApiGatewayHost,
		//AllowCredentials: true,
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, HEAD, PUT, PATCH, POST, DELETE",
	}))
	f.Use(log.RequestLogger(h.logger))

	f.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("film service healthy")
	})

	f.Listen(fmt.Sprintf(":%s", h.conf.Application.FilmServicePort))
}
