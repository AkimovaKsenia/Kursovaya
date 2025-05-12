package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"kino/internal/shared/entities"
	"kino/internal/shared/log"
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
	return nil
}

func (h *Handler) GetAllCinemaHallsByID(c *fiber.Ctx) error {
	return nil
}

func (h *Handler) GetCinemaByID(c *fiber.Ctx) error {
	return nil
}

func (h *Handler) UpdateCinema(c *fiber.Ctx) error {
	return nil
}

func (h *Handler) DeleteCinema(c *fiber.Ctx) error {
	return nil
}
