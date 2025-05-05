package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"kino/internal/shared/entities"
	"kino/internal/shared/log"
	"kino/pkg/util"
)

// Login
// @Tags         Auth
// @Summary      Аутентификация пользователя
// @Description  Вход пользователя в систему с выдачей jwt и роли
// @Accept       json
// @Produce      json
// @Param        request body entities.LoginUserRequest true "Данные для входа"
// @Success      200 {object} entities.LoginUserResponse "Успешная аутентификация"
// @Failure      400 {object} entities.Error "Некорректные данные для входа"
// @Failure      500 {object} entities.Error "Ошибка на стороне сервера"
// @Router       /login [post]
func (h *Handler) Login(c *fiber.Ctx) error {
	var user entities.LoginUserRequest
	if err := c.BodyParser(&user); err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	h.logger.Debug().Msg("call h.repository.DB.DBUserGetByEmail")
	u, err := h.repository.DB.GetUserByEmail(user.Email)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(fmt.Sprintf("error getting user by email: %s", err.Error()))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("error getting user by email: %s", err.Error())})
	}

	h.logger.Debug().Msg("call util.CheckPassword")
	fmt.Println(user, u)
	err = util.CheckPassword(user.Password, u.Password)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg("wrong data")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "wrong data"})
	}

	TokenExpiration := 100
	h.logger.Debug().Msg("call pkg.GenerateAccessToken")
	token, err := h.GenerateToken(u.ID, TokenExpiration)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	res := entities.LoginUserResponse{
		Token: token,
		Role:  u.Role,
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(res)
}
