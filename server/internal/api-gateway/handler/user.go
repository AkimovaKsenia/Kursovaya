package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
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

// Register
// @Tags         User
// @Summary      Создание нового пользователя
// @Description  Создание пользователя с переданными данными (только для админа)
// @Accept       json
// @Produce      json
// @Param        request body entities.CreateUser true "Данные пользователя"
// @Success      200 {object} entities.ID "Пользователь создан"
// @Failure      400 {object} entities.Error "Некорректные данные для входа"
// @Failure      403 {object} entities.Error "Недостаточно прав"
// @Failure      500 {object} entities.Error "Ошибка на стороне сервера"
// @Router       /auth/register [post]
// @Security ApiKeyAuth
func (h *Handler) Register(c *fiber.Ctx) error {
	h.logger.Debug().Caller().Msg("body parse")
	var f entities.User
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

	if f.Surname == "" {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg("empty field 'surname'")
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "empty field 'surname'"})
	}

	if f.Email == "" {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg("empty field 'email'")
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "empty field 'email'"})
	}

	if f.Password == "" {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg("empty field 'password'")
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "empty field 'password'"})
	}

	if f.RoleID <= 0 {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Int("id", f.RoleID).Msg("invalid field 'role_id'")
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "invalid field 'role_id'"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(f.Password), bcrypt.DefaultCost)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(fmt.Sprintf("error hashing password: %s", err.Error()))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("error hashing password: %s", err.Error())})
	}
	f.Password = string(hashedPassword)

	h.logger.Debug().Msg("call h.repository.DB.CreateUser")
	id, err := h.repository.DB.CreateUser(&f)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(fmt.Sprintf("error creating user: %s", err.Error()))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("error creating user: %s", err.Error())})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("successful creating user")
	return c.Status(fiber.StatusOK).JSON(entities.ID{ID: id})
}

// GetUsers
// @Tags         User
// @Summary      Получение всех пользователей
// @Accept       json
// @Produce      json
// @Success      200 {object} []entities.GetUser "Пользователи получены"
// @Failure      400 {object} entities.Error "Некорректные данные для входа"
// @Failure      403 {object} entities.Error "Недостаточно прав"
// @Failure      500 {object} entities.Error "Ошибка на стороне сервера"
// @Router       /auth/user [get]
// @Security ApiKeyAuth
func (h *Handler) GetUsers(c *fiber.Ctx) error {
	h.logger.Debug().Msg("calling h.repository.DB.GetAllUsers")
	users, err := h.repository.DB.GetAllUsers()
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(fmt.Sprintf("error getting all users: %s", err.Error()))
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: fmt.Sprintf("error getting all users: %s", err.Error())})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("successful getting all users")
	return c.Status(fiber.StatusOK).JSON(users)
}

// GetUserRoles
// @Tags         User
// @Summary      Получение всех ролей
// @Accept       json
// @Produce      json
// @Success      200 {object} []entities.Role "Роли получены"
// @Failure      400 {object} entities.Error "Некорректные данные для входа"
// @Failure      403 {object} entities.Error "Недостаточно прав"
// @Failure      500 {object} entities.Error "Ошибка на стороне сервера"
// @Router       /auth/user/role [get]
// @Security ApiKeyAuth
func (h *Handler) GetUserRoles(c *fiber.Ctx) error {
	roles := []entities.Role{
		{
			ID:   1,
			Name: "Работник",
		},
		{
			ID:   2,
			Name: "Администратор",
		},
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusNotFound).JSON(roles)
}
