package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"kino/internal/shared/entities"
	"kino/internal/shared/log"
	"strings"
)

// CreateFilm
// @Tags         Film
// @Summary      Создание фильма
// @Accept multipart/form-data
// @Produce      json
// @Param name formData string true "Название" example("Оно")
// @Param description formData string true "Описание" example("Сука клоун")
// @Param film_photo formData file false "Фото фильма"
// @Param cast_list formData []string true "Список актеров" example("Райан Рейнольдс, Морена Баккарин")
// @Param film_studio_id formData int true "ID киностудии" example("1")
// @Param duration_in_min formData int true "Продолжительность фильма в минутах" example("123")
// @Param director_ids formData []int true "ID режиссеров"
// @Param operator_ids formData []int true "ID операторов"
// @Param genre_ids formData []int true "ID жанров"
// @Success      200 {object} entities.ID "Фильм успешно создан"
// @Failure      400 {object} entities.Error "Некорректный запрос"
// @Failure      401 {object} entities.Error "Пользователь не авторизован"
// @Failure      403 {object} entities.Error "Недостаточно прав для запроса"
// @Failure      500 {object} entities.Error "Ошибка на стороне сервера"
// @Router       /auth/film [post]
// @Security ApiKeyAuth
func (h *Handler) CreateFilm(c *fiber.Ctx) error {
	requestURL := fmt.Sprintf("%s/%s", h.conf.Application.FilmServiceHost, strings.TrimPrefix(c.OriginalURL(), "/auth/"))
	c.Locals("request_url", requestURL)
	c.Locals("request_method", fiber.MethodPost)
	c.Locals("content_type", "multipart/form-data")

	h.logger.Debug().Msg("call h.Redirect")
	responseBody, err := h.Redirect(c)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(fmt.Sprintf("error sending request to film service: %s", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{Error: fmt.Sprintf("error sending request to film service: %s", err.Error())})
	}

	return c.Send(responseBody)
}

// GetFilmByID
// @Tags         Film
// @Summary      Получение фильма по ID
// @Accept json
// @Produce      json
// @Param id path int true "ID фильма" example(1)
// @Success      200 {object} entities.FilmFull "Фильм по ID успешно получен"
// @Failure      400 {object} entities.Error "Некорректный запрос"
// @Failure      500 {object} entities.Error "Ошибка на стороне сервера"
// @Router       /auth/film/id/{id} [get]
// @Security ApiKeyAuth
func (h *Handler) GetFilmByID(c *fiber.Ctx) error {
	requestURL := fmt.Sprintf("%s/%s", h.conf.Application.FilmServiceHost, strings.TrimPrefix(c.OriginalURL(), "/auth/"))
	c.Locals("request_url", requestURL)
	c.Locals("request_method", fiber.MethodGet)

	h.logger.Debug().Msg("call h.Redirect")
	responseBody, err := h.Redirect(c)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(fmt.Sprintf("error sending request to film service: %s", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{Error: fmt.Sprintf("error sending request to film service: %s", err.Error())})
	}

	return c.Send(responseBody)
}

// GetAllFilms
// @Tags         Film
// @Summary      Получение всех фильмов
// @Accept json
// @Produce      json
// @Success      200 {object} []entities.FilmFull "Все фильмы успешно получены"
// @Failure      400 {object} entities.Error "Некорректный запрос"
// @Failure      500 {object} entities.Error "Ошибка на стороне сервера"
// @Router       /auth/film [get]
// @Security ApiKeyAuth
func (h *Handler) GetAllFilms(c *fiber.Ctx) error {
	requestURL := fmt.Sprintf("%s/%s", h.conf.Application.FilmServiceHost, strings.TrimPrefix(c.OriginalURL(), "/auth/"))
	c.Locals("request_url", requestURL)
	c.Locals("request_method", fiber.MethodGet)

	h.logger.Debug().Msg("call h.Redirect")
	responseBody, err := h.Redirect(c)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(fmt.Sprintf("error sending request to film service: %s", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{Error: fmt.Sprintf("error sending request to film service: %s", err.Error())})
	}

	return c.Send(responseBody)
}

// GetAllGenres
// @Tags         Film
// @Summary      Получение всех жанров
// @Accept       json
// @Produce      json
// @Success      200  {array}   entities.Genre  "Список всех жанров"
// @Failure      500  {object}  entities.Error  "Ошибка на стороне сервера"
// @Router       /auth/genres [get]
// @Security ApiKeyAuth
func (h *Handler) GetAllGenres(c *fiber.Ctx) error {
	requestURL := fmt.Sprintf("%s/%s", h.conf.Application.FilmServiceHost, strings.TrimPrefix(c.OriginalURL(), "/auth/"))
	c.Locals("request_url", requestURL)
	c.Locals("request_method", fiber.MethodGet)

	h.logger.Debug().Msg("call h.Redirect")
	responseBody, err := h.Redirect(c)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(fmt.Sprintf("error sending request to film service: %s", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{Error: fmt.Sprintf("error sending request to film service: %s", err.Error())})
	}

	return c.Send(responseBody)
}

// GetAllOperators
// @Tags         Film
// @Summary      Получение всех операторов
// @Accept       json
// @Produce      json
// @Success      200  {array}   entities.Operator  "Список всех операторов"
// @Failure      500  {object}  entities.Error     "Ошибка на стороне сервера"
// @Router       /auth/operators [get]
// @Security ApiKeyAuth
func (h *Handler) GetAllOperators(c *fiber.Ctx) error {
	requestURL := fmt.Sprintf("%s/%s", h.conf.Application.FilmServiceHost, strings.TrimPrefix(c.OriginalURL(), "/auth/"))
	c.Locals("request_url", requestURL)
	c.Locals("request_method", fiber.MethodGet)

	h.logger.Debug().Msg("call h.Redirect")
	responseBody, err := h.Redirect(c)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(fmt.Sprintf("error sending request to film service: %s", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{Error: fmt.Sprintf("error sending request to film service: %s", err.Error())})
	}

	return c.Send(responseBody)
}

// GetAllDirectors
// @Tags         Film
// @Summary      Получение всех режиссеров
// @Accept       json
// @Produce      json
// @Success      200  {array}   entities.Director  "Список всех режиссеров"
// @Failure      500  {object}  entities.Error     "Ошибка на стороне сервера"
// @Router       /auth/directors [get]
// @Security ApiKeyAuth
func (h *Handler) GetAllDirectors(c *fiber.Ctx) error {
	requestURL := fmt.Sprintf("%s/%s", h.conf.Application.FilmServiceHost, strings.TrimPrefix(c.OriginalURL(), "/auth/"))
	c.Locals("request_url", requestURL)
	c.Locals("request_method", fiber.MethodGet)

	h.logger.Debug().Msg("call h.Redirect")
	responseBody, err := h.Redirect(c)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(fmt.Sprintf("error sending request to film service: %s", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{Error: fmt.Sprintf("error sending request to film service: %s", err.Error())})
	}

	return c.Send(responseBody)
}

// GetAllFilmStudios
// @Tags         Film
// @Summary      Получение всех киностудий
// @Accept       json
// @Produce      json
// @Success      200  {array}   entities.FilmStudio  "Список всех киностудий"
// @Failure      500  {object}  entities.Error       "Ошибка на стороне сервера"
// @Router       /auth/film-studios [get]
// @Security ApiKeyAuth
func (h *Handler) GetAllFilmStudios(c *fiber.Ctx) error {
	requestURL := fmt.Sprintf("%s/%s", h.conf.Application.FilmServiceHost, strings.TrimPrefix(c.OriginalURL(), "/auth/"))
	c.Locals("request_url", requestURL)
	c.Locals("request_method", fiber.MethodGet)

	h.logger.Debug().Msg("call h.Redirect")
	responseBody, err := h.Redirect(c)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(fmt.Sprintf("error sending request to film service: %s", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{Error: fmt.Sprintf("error sending request to film service: %s", err.Error())})
	}

	return c.Send(responseBody)
}
