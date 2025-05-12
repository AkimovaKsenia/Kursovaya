package handler

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"kino/internal/shared/entities"
	"kino/internal/shared/log"
	"regexp"
	"strings"
)

func (h *Handler) GetAllCinemaConditions(c *fiber.Ctx) error {
	h.logger.Debug().Msg("calling h.repository.DB.GetAllCinemaConditions")
	cinemaConditions, err := h.repository.DB.GetAllCinemaConditions()
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(fmt.Sprintf("error getting all cinema conditions: %s", err.Error()))
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: fmt.Sprintf("error getting all cinema conditions: %s", err.Error())})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("successful getting all cinema conditions")
	return c.Status(fiber.StatusOK).JSON(cinemaConditions)
}

func (h *Handler) GetAllCinemaCategories(c *fiber.Ctx) error {
	h.logger.Debug().Msg("calling h.repository.DB.GetAllCinemaCategories")
	cinemaCategories, err := h.repository.DB.GetAllCinemaCategories()
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(fmt.Sprintf("error getting all cinema categories: %s", err.Error()))
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: fmt.Sprintf("error getting all cinema categories: %s", err.Error())})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("successful getting all cinema categories")
	return c.Status(fiber.StatusOK).JSON(cinemaCategories)
}

func (h *Handler) GetAllCinemaHallTypes(c *fiber.Ctx) error {
	h.logger.Debug().Msg("calling h.repository.DB.GetAllCinemaHallTypes")
	cinemaHallTypes, err := h.repository.DB.GetAllCinemaHallTypes()
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(fmt.Sprintf("error getting all cinema hall types: %s", err.Error()))
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: fmt.Sprintf("error getting all cinema hall types: %s", err.Error())})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("successful getting all cinema hall types")
	return c.Status(fiber.StatusOK).JSON(cinemaHallTypes)
}

func (h *Handler) CreateCinema(c *fiber.Ctx) error {
	return nil
}

func (h *Handler) GetAllCinemasAddressName(c *fiber.Ctx) error {
	h.logger.Debug().Msg("calling h.repository.DB.GetAllCinemasAddressName")
	cinemasAddressName, err := h.repository.DB.GetAllCinemasAddressName()
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(fmt.Sprintf("error getting all cinema address and name: %s", err.Error()))
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: fmt.Sprintf("error getting all cinema address and name: %s", err.Error())})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("successful getting all cinema address and name")
	return c.Status(fiber.StatusOK).JSON(cinemasAddressName)
}

func (h *Handler) GetAllCinemaHallsByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(fmt.Sprintf("invalid film ID: %v", err))
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "Invalid film ID"})
	}

	h.logger.Debug().Msg("calling h.repository.DB.GetAllCinemaHallsByID")
	cinemaHallTypes, err := h.repository.DB.GetAllCinemaHallsByID(id)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(fmt.Sprintf("error getting all cinema halls by id: %s", err.Error()))
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: fmt.Sprintf("error getting all cinema halls by id: %s", err.Error())})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("successful getting all cinema halls by id")
	return c.Status(fiber.StatusOK).JSON(cinemaHallTypes)
}

func (h *Handler) GetCinemaByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(fmt.Sprintf("invalid film ID: %v", err))
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "Invalid film ID"})
	}

	h.logger.Debug().Msg("calling h.repository.DB.GetCinemaByID")
	cinema, err := h.repository.DB.GetCinemaByID(id)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(fmt.Sprintf("error getting cinema by id: %s", err.Error()))
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: fmt.Sprintf("error getting cinema by id: %s", err.Error())})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("successful getting cinema by id")
	return c.Status(fiber.StatusOK).JSON(cinema)
}

func (h *Handler) UpdateCinema(c *fiber.Ctx) error {
	return nil
}

func (h *Handler) DeleteCinema(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(fmt.Sprintf("invalid cinema ID: %v", err))
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "Invalid film ID"})
	}

	cinemaDB, err := h.repository.DB.GetCinemaByID(id)
	if err != nil {
		status := fiber.StatusInternalServerError
		if strings.Contains(err.Error(), "not found") {
			status = fiber.StatusNotFound
		}

		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: status})
		logEvent.Msg(fmt.Sprintf("error deleting cinema: error getting cinema by id: %v", err))
		return c.Status(status).JSON(entities.Error{Error: err.Error()})
	}

	re := regexp.MustCompile("^(https?|ftp):\\/\\/[^\\s/$.?#].[^\\s]*$")

	if !re.MatchString(cinemaDB.Photo) && cinemaDB.Photo != "" {
		err = h.repository.S3.RemoveObject(context.Background(), "cinema-media", cinemaDB.Photo)
		if err != nil {
			logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
			logEvent.Msg(fmt.Sprintf("error removing file %s from minio: %s", cinemaDB.Photo, err.Error()))
			return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: fmt.Sprintf("error removing file %s from minio: %s", cinemaDB.Photo, err.Error())})
		}
	}

	if err := h.repository.DB.DeleteFilm(id); err != nil {
		status := fiber.StatusInternalServerError
		if strings.Contains(err.Error(), "not found") {
			status = fiber.StatusNotFound
		}

		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: status})
		logEvent.Msg(fmt.Sprintf("error deleting film: %v", err))
		return c.Status(status).JSON(entities.Error{Error: err.Error()})
	}

	return c.JSON(entities.ID{ID: id})
}
