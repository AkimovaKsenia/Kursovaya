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
		AllowOrigins:     h.conf.Application.ApiGatewayHost,
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, HEAD, PUT, PATCH, POST, DELETE",
	}))
	f.Use(log.RequestLogger(h.logger))

	f.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("cinema service healthy")
	})

	f.Get("/cinema/conditions", h.GetAllCinemaConditions)
	f.Get("/cinema/categories", h.GetAllCinemaCategories)
	f.Get("/cinema/hall/types", h.GetAllCinemaHallTypes)

	f.Post("/cinema", h.CreateCinema)
	f.Get("/cinema/address_name", h.GetAllCinemasAddressName)
	f.Get("/cinema/halls/:id", h.GetAllCinemaHallsByID)
	f.Get("/cinema/id/:id", h.GetCinemaByID)
	f.Put("/cinema", h.UpdateCinema)
	f.Delete("/cinema/:id", h.DeleteCinema)

	h.logger.Info().Msg(fmt.Sprintf("start cinema service on port %s", h.conf.Application.ApiGatewayPort))
	f.Listen(fmt.Sprintf(":%s", h.conf.Application.CinemaServicePort))
}
