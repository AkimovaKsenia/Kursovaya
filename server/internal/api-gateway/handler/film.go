package handler

import (
	"context"
	"fmt"
	"kino/internal/shared/entities"
	"kino/internal/shared/log"
	"kino/pkg/util"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
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
	userId := c.Locals("id").(int)
	uRole, err := h.repository.DB.GetUserRoleById(userId)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(fmt.Sprintf("error getting user role: %s", err.Error()))
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: fmt.Sprintf("error getting user role: %s", err.Error())})
	}

	if uRole.Role != "admin" {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusForbidden})
		logEvent.Str("user_role", uRole.Role).Msg("there are not enough rights for this action")
		return c.Status(fiber.StatusForbidden).JSON(entities.Error{Error: "there are not enough rights for this action"})
	}

	var f entities.CreateFilm
	if err := c.BodyParser(&f); err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: err.Error()})
	}

	file, err := c.FormFile("photo")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Photo file is required",
		})
	}
	filePath := fmt.Sprintf("./tmp/%s", file.Filename)

	c.SaveFile(file, filePath)
	fileName := util.GenerateRandomFileName(filepath.Ext(file.Filename))

	err = h.repository.S3.FPutObject(context.Background(), "film-media", fileName, filePath, file.Header.Get("Content-Type"))
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(fmt.Sprintf("error creating file %s in minio: %s", fileName, err.Error()))
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: fmt.Sprintf("error creating file %s in minio: %s", file.Filename, err.Error())})
	}
	err = os.Remove(filePath)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(fmt.Sprintf("error removing file %s: %s", file.Filename, err.Error()))
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: fmt.Sprintf("error removing file %s: %s", file.Filename, err.Error())})
	}

	f.Photo = fileName

	h.logger.Debug().Msg("call h.repository.DB.DBUserGetByEmail")
	h.repository.DB.CreateFilm(&f)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(fmt.Sprintf("error creating film: %s", err.Error()))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("error creating film: %s", err.Error())})
	}

	id := 1
	fmt.Println(f)
	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg(fmt.Sprintf("successful creating university with id=%v", id))
	return c.Status(fiber.StatusOK).JSON(entities.ID{ID: id})
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
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: err.Error()})
	}
	h.logger.Debug().Msg("calling h.repository.DB.GetFilmByID")
	film, err := h.repository.DB.GetFilmByID(id)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(fmt.Sprintf("error getting film with id=%v: %s", id, err.Error()))
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: fmt.Sprintf("error getting film by id=%v: %s", id, err.Error())})
	}

	re := regexp.MustCompile("^(https?|ftp):\\/\\/[^\\s/$.?#].[^\\s]*$")

	if !re.MatchString(film.Photo) {
		url, _ := h.repository.S3.PresignedGetObject(context.Background(), "film-media", film.Photo, 7*24*time.Hour)
		film.Photo = url.String()
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg(fmt.Sprintf("successful getting film with id=%v", id))
	return c.Status(fiber.StatusOK).JSON(film)
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
	h.logger.Debug().Msg("calling h.repository.DB.GetFilmByID")
	films, err := h.repository.DB.GetAllFilms()
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(fmt.Sprintf("error getting all films: %s", err.Error()))
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: fmt.Sprintf("error getting all films: %s", err.Error())})
	}

	re := regexp.MustCompile("^(https?|ftp):\\/\\/[^\\s/$.?#].[^\\s]*$")

	for i := range films {
		if !re.MatchString(films[i].Photo) {
			url, _ := h.repository.S3.PresignedGetObject(context.Background(), "film-media", films[i].Photo, 7*24*time.Hour)
			films[i].Photo = url.String()
		}
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("successful getting all films")
	return c.Status(fiber.StatusOK).JSON(films)
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
	genres, err := h.repository.DB.GetAllGenres()
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(fmt.Sprintf("error getting genres: %s", err.Error()))

		return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{Error: "Failed to get genres list"})
	}
	return c.JSON(genres)
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
	operators, err := h.repository.DB.GetAllOperators()
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(fmt.Sprintf("error getting operators: %s", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{Error: "Failed to get operators list"})
	}
	return c.JSON(operators)
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
	directors, err := h.repository.DB.GetAllDirectors()
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{
			Level:  "Error",
			Method: c.Method(),
			Url:    c.OriginalURL(),
			Status: fiber.StatusInternalServerError,
		})
		logEvent.Msg(fmt.Sprintf("error getting directors: %s", err.Error()))

		return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{
			Error: "Failed to get directors list",
		})
	}
	return c.JSON(directors)
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
