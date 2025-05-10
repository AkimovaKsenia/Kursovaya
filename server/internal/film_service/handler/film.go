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
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateFilm(c *fiber.Ctx) error {
	h.logger.Debug().Caller().Msg("body parse")
	var f entities.CreateFilm
	if err := c.BodyParser(&f); err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: err.Error()})
	}

	if f.Name == "" {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg("empty field 'name'")
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "empty field 'name'"})
	}

	if f.Description == "" {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg("empty field 'description'")
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "empty field 'description'"})
	}

	if len(f.CastList) == 0 {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg("invalid field 'cast_list'")
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "invalid field 'cast_list'"})
	}

	if f.FilmStudioID <= 0 {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg("invalid field 'film_studio_id'")
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "invalid field 'film_studio_id'"})
	}

	if f.DurationInMin <= 0 {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg("invalid field 'duration_in_min'")
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "invalid field 'duration_in_min'"})
	}

	if len(f.DirectorIDs) == 0 {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg("empty field 'duration_in_min'")
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "empty field 'duration_in_min'"})
	}

	if len(f.GenreIDs) == 0 {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg("empty field 'director_ids'")
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "empty field 'director_ids'"})
	}

	file, err := c.FormFile("film_photo")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Photo file is required",
		})
	}
	filePath := fmt.Sprintf("./tmp/%s", file.Filename)

	h.logger.Debug().Caller().Msg("save file")
	c.SaveFile(file, filePath)
	fileName := util.GenerateRandomFileName(filepath.Ext(file.Filename))

	h.logger.Debug().Caller().Msg("call h.repository.S3.FPutObject")
	err = h.repository.S3.FPutObject(context.Background(), "film-media", fileName, filePath, file.Header.Get("Content-Type"))
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(fmt.Sprintf("error creating file %s in minio: %s", fileName, err.Error()))
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: fmt.Sprintf("error creating file %s in minio: %s", file.Filename, err.Error())})
	}
	h.logger.Debug().Caller().Msg("remove file")
	err = os.Remove(filePath)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(fmt.Sprintf("error removing file %s: %s", file.Filename, err.Error()))
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: fmt.Sprintf("error removing file %s: %s", file.Filename, err.Error())})
	}

	f.Photo = fileName

	h.logger.Debug().Msg("call h.repository.DB.DBUserGetByEmail")
	id, err := h.repository.DB.CreateFilm(&f)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(fmt.Sprintf("error creating film: %s", err.Error()))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("error creating film: %s", err.Error())})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg(fmt.Sprintf("successful creating university with id=%v", id))
	return c.Status(fiber.StatusOK).JSON(entities.ID{ID: id})
}

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

func (h *Handler) GetAllGenres(c *fiber.Ctx) error {
	genres, err := h.repository.DB.GetAllGenres()
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(fmt.Sprintf("error getting genres: %s", err.Error()))

		return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{Error: "Failed to get genres list"})
	}
	return c.JSON(genres)
}

func (h *Handler) GetAllOperators(c *fiber.Ctx) error {
	operators, err := h.repository.DB.GetAllOperators()
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(fmt.Sprintf("error getting operators: %s", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{Error: "Failed to get operators list"})
	}
	return c.JSON(operators)
}

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

func (h *Handler) GetAllFilmStudios(c *fiber.Ctx) error {
	studios, err := h.repository.DB.GetAllFilmStudios()
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{
			Level:  "Error",
			Method: c.Method(),
			Url:    c.OriginalURL(),
			Status: fiber.StatusInternalServerError,
		})
		logEvent.Msg(fmt.Sprintf("error getting film studios: %s", err.Error()))

		return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{
			Error: "Failed to get film studios list",
		})
	}
	return c.JSON(studios)
}
