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
		authGroup.Get("/genres", h.GetAllGenres)
		authGroup.Get("/operators", h.GetAllOperators)
		authGroup.Get("/directors", h.GetAllDirectors)
		authGroup.Get("/film-studios", h.GetAllFilmStudios)

		authGroup.Post("/film", h.CreateFilm)
		authGroup.Get("/film", h.GetAllFilms)
		authGroup.Get("/film/id/:id", h.GetFilmByID)
	}

	h.logger.Info().Msg(fmt.Sprintf("start api-gateway on port %s", h.conf.Application.ApiGatewayPort))
	f.Listen(fmt.Sprintf(":%s", h.conf.Application.ApiGatewayPort))
}

func (h *Handler) Redirect(c *fiber.Ctx) ([]byte, error) {
	requestURL := c.Locals("request_url")
	if requestURL == nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(fmt.Sprintf("error creating request to film service: %s", "empty request url"))
		return nil, fmt.Errorf(fmt.Sprintf("error creating request to film service: %s", "empty request url"))
	}

	userID := c.Locals("id").(int)
	requestMethod := c.Locals("request_method")
	if requestURL == nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(fmt.Sprintf("error creating request to film service: %s", "empty request url"))
		return nil, fmt.Errorf(fmt.Sprintf("error creating request to film service: %s", "empty request url"))
	}

	h.logger.Debug().Caller().Msg("create new request")
	req, err := http.NewRequestWithContext(
		c.Context(),
		requestMethod.(string),
		requestURL.(string),
		bytes.NewReader(c.Request().Body()),
	)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(fmt.Sprintf("error creating request to film service: %s", err.Error()))
		return nil, fmt.Errorf(fmt.Sprintf("error creating request to film service: %s", err.Error()))
	}

	copyHeaders(c, req)

	req.Header.Set("X-User-ID", fmt.Sprintf("%d", userID))

	h.logger.Debug().Caller().Msg("sending request to service")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{
			Level:  "Error",
			Method: c.Method(),
			Url:    c.OriginalURL(),
			Status: fiber.StatusBadGateway,
		})
		logEvent.Msg(fmt.Sprintf("error calling film service: %s", err.Error()))
		return nil, fmt.Errorf(fmt.Sprintf("error calling film service: %s", err.Error()))
	}
	defer resp.Body.Close()

	h.logger.Debug().Caller().Msg("reading response body")
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{
			Level:  "Error",
			Method: c.Method(),
			Url:    c.OriginalURL(),
			Status: fiber.StatusInternalServerError,
		})
		logEvent.Msg(fmt.Sprintf("error reading film service response: %s", err.Error()))
		return nil, fmt.Errorf(fmt.Sprintf("error calling film service: %s", err.Error()))
	}

	c.Status(resp.StatusCode)
	for k, v := range resp.Header {
		for _, h := range v {
			c.Set(k, h)
		}
	}

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
