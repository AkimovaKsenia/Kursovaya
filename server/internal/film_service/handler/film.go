package handler

import (
	"fmt"
	"kino/internal/shared/entities"
	"kino/internal/shared/log"

	"github.com/gofiber/fiber/v2"
)

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
