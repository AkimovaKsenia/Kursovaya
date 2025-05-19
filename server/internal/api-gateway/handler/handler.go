package handler

import (
	"bytes"
	"fmt"
	"io"
	_ "kino/docs"
	"kino/internal/shared/config"
	"kino/internal/shared/log"
	"kino/internal/shared/repository"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rs/zerolog"
	fiberSwagger "github.com/swaggo/fiber-swagger"
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
		AllowOrigins: "*",
		//AllowCredentials: true,
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, HEAD, PUT, PATCH, POST, DELETE",
	}))
	f.Use(log.RequestLogger(h.logger))

	f.Get("/swagger/*", fiberSwagger.WrapHandler)
	f.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("api-gateway healthy")
	})

	f.Post("/login", h.Login)

	authGroup := f.Group("/auth")
	authGroup.Use(func(c *fiber.Ctx) error {
		return h.WithJWTAuth(c)
	})
	{
		// FILM
		{
			authGroup.Get("/genres", h.GetAllGenres)
			authGroup.Get("/operators", h.GetAllOperators)
			authGroup.Get("/directors", h.GetAllDirectors)
			authGroup.Get("/film-studios", h.GetAllFilmStudios)

			authGroup.Post("/film", h.CreateFilm)
			authGroup.Get("/film", h.GetAllFilms)
			authGroup.Get("/film/id/:id", h.GetFilmByID)
			authGroup.Put("/film", h.UpdateFilm)
			authGroup.Delete("/film/:id", h.DeleteFilm)
		}

		// CINEMA
		{
			authGroup.Get("/cinema/conditions", h.GetAllCinemaConditions)
			authGroup.Get("/cinema/categories", h.GetAllCinemaCategories)
			authGroup.Get("/cinema/hall/types", h.GetAllCinemaHallTypes)

			authGroup.Post("/cinema", h.CreateCinema)
			authGroup.Post("/cinema_hall", h.CreateCinemaHall)
			authGroup.Get("/cinema/address_name", h.GetAllCinemasAddressName)
			authGroup.Get("/cinema/halls/:id", h.GetAllCinemaHallsByID)
			authGroup.Get("/cinema/id/:id", h.GetCinemaByID)
			authGroup.Put("/cinema", h.UpdateCinema)
			authGroup.Put("/cinema_hall", h.UpdateCinemaHall)
			authGroup.Delete("/cinema/:id", h.DeleteCinema)
			authGroup.Delete("/cinema_hall/:id", h.DeleteCinemaHall)
		}

	}

	h.logger.Info().Msg(fmt.Sprintf("start api-gateway on port %s", h.conf.Application.ApiGatewayPort))
	f.Listen(fmt.Sprintf(":%s", h.conf.Application.ApiGatewayPort))
}

func (h *Handler) Redirect(c *fiber.Ctx) ([]byte, error) {
	requestURL := c.Locals("request_url")
	requestMethod := c.Method()
	userID := c.Locals("id")
	if requestURL == nil || requestMethod == "" || userID == nil {
		return nil, fmt.Errorf("error creating request to service: missing required locals (request_url, request_method or id)")
	}

	contentType := c.Get("Content-Type")
	if contentType == "" {
		contentType = "application/json"
	}

	h.logger.Debug().Caller().Msg("create new request")
	req, err := http.NewRequestWithContext(
		c.Context(),
		requestMethod,
		requestURL.(string),
		bytes.NewReader(c.BodyRaw()),
	)
	if err != nil {
		return nil, fmt.Errorf("error creating request to service: %v", err)
	}

	copyHeaders(c, req)
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("X-User-ID", fmt.Sprintf("%v", userID))

	h.logger.Debug().Caller().Msg("sending request to service")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request to service: %v", err)
	}
	defer resp.Body.Close()

	h.logger.Debug().Caller().Msg("reading response body")
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response from service: %v", err)
	}

	c.Set("Content-Type", "application/json")
	c.Status(resp.StatusCode)

	return responseBody, nil
}

func copyHeaders(c *fiber.Ctx, req *http.Request) {
	for k, values := range c.GetReqHeaders() {
		if strings.EqualFold(k, "Content-Length") || strings.EqualFold(k, "Host") {
			continue
		}
		for _, v := range values {
			req.Header.Add(k, v)
		}
	}
}
