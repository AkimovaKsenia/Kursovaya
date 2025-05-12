package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"kino/internal/shared/entities"
	"kino/internal/shared/log"
)

// GetAllCinemaConditions
// @Tags         Film
// @Summary      Получение всех жанров
// @Accept       json
// @Produce      json
// @Success      200  {array}   entities.CinemaCondition  "Список всех жанров"
// @Failure      500  {object}  entities.Error  "Ошибка на стороне сервера"
// @Router       /auth/cinema/conditions [get]
// @Security ApiKeyAuth
func (h *Handler) GetAllCinemaConditions(c *fiber.Ctx) error {
	h.logger.Debug().Msg("call h.Redirect")
	responseBody, err := h.Redirect(c)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(fmt.Sprintf("error sending request to cinema service: %s", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{Error: fmt.Sprintf("error sending request to film service: %s", err.Error())})
	}

	return c.Send(responseBody)
}

func (h *Handler) GetAllCinemaCategories(c *fiber.Ctx) error {
	h.logger.Debug().Msg("call h.Redirect")
	responseBody, err := h.Redirect(c)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(fmt.Sprintf("error sending request to cinema service: %s", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{Error: fmt.Sprintf("error sending request to film service: %s", err.Error())})
	}

	return c.Send(responseBody)
}

func (h *Handler) GetAllCinemaHallTypes(c *fiber.Ctx) error {
	h.logger.Debug().Msg("call h.Redirect")
	responseBody, err := h.Redirect(c)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(fmt.Sprintf("error sending request to cinema service: %s", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{Error: fmt.Sprintf("error sending request to film service: %s", err.Error())})
	}

	return c.Send(responseBody)
}

func (h *Handler) CreateCinema(c *fiber.Ctx) error {
	h.logger.Debug().Msg("call h.Redirect")
	responseBody, err := h.Redirect(c)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(fmt.Sprintf("error sending request to cinema service: %s", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{Error: fmt.Sprintf("error sending request to film service: %s", err.Error())})
	}

	return c.Send(responseBody)
}

func (h *Handler) GetAllCinemasAddressName(c *fiber.Ctx) error {
	h.logger.Debug().Msg("call h.Redirect")
	responseBody, err := h.Redirect(c)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(fmt.Sprintf("error sending request to cinema service: %s", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{Error: fmt.Sprintf("error sending request to film service: %s", err.Error())})
	}

	return c.Send(responseBody)
}

func (h *Handler) GetAllCinemaHallsByID(c *fiber.Ctx) error {
	h.logger.Debug().Msg("call h.Redirect")
	responseBody, err := h.Redirect(c)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(fmt.Sprintf("error sending request to cinema service: %s", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{Error: fmt.Sprintf("error sending request to film service: %s", err.Error())})
	}

	return c.Send(responseBody)
}

func (h *Handler) GetCinemaByID(c *fiber.Ctx) error {
	h.logger.Debug().Msg("call h.Redirect")
	responseBody, err := h.Redirect(c)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(fmt.Sprintf("error sending request to cinema service: %s", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{Error: fmt.Sprintf("error sending request to film service: %s", err.Error())})
	}

	return c.Send(responseBody)
}

func (h *Handler) UpdateCinema(c *fiber.Ctx) error {
	h.logger.Debug().Msg("call h.Redirect")
	responseBody, err := h.Redirect(c)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(fmt.Sprintf("error sending request to cinema service: %s", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{Error: fmt.Sprintf("error sending request to film service: %s", err.Error())})
	}

	return c.Send(responseBody)
}

func (h *Handler) DeleteCinema(c *fiber.Ctx) error {
	h.logger.Debug().Msg("call h.Redirect")
	responseBody, err := h.Redirect(c)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(fmt.Sprintf("error sending request to cinema service: %s", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{Error: fmt.Sprintf("error sending request to film service: %s", err.Error())})
	}

	return c.Send(responseBody)
}
