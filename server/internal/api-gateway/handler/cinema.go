package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"kino/internal/shared/entities"
	"kino/internal/shared/log"
	"strings"
)

// GetAllCinemaConditions
// @Tags         Cinema
// @Summary      Получение всех состояний кинотеатра
// @Accept       json
// @Produce      json
// @Success      200  {array}   []entities.CinemaCondition  "Список всех состояний кинотеатра"
// @Failure      500  {object}  entities.Error  "Ошибка на стороне сервера"
// @Router       /auth/cinema/conditions [get]
// @Security ApiKeyAuth
func (h *Handler) GetAllCinemaConditions(c *fiber.Ctx) error {
	requestURL := fmt.Sprintf("%s/%s", h.conf.Application.CinemaServiceHost, strings.TrimPrefix(c.OriginalURL(), "/auth/"))
	c.Locals("request_url", requestURL)
	h.logger.Debug().Msg("call h.Redirect")
	responseBody, err := h.Redirect(c)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(fmt.Sprintf("error sending request to cinema service: %s", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{Error: fmt.Sprintf("error sending request to film service: %s", err.Error())})
	}

	return c.Send(responseBody)
}

// GetAllCinemaCategories
// @Tags         Cinema
// @Summary      Получение всех категорий кинотеатра
// @Accept       json
// @Produce      json
// @Success      200  {array}   []entities.CinemaCategory  "Список всех категорий кинотеатра"
// @Failure      500  {object}  entities.Error  "Ошибка на стороне сервера"
// @Router       /auth/cinema/categories [get]
// @Security ApiKeyAuth
func (h *Handler) GetAllCinemaCategories(c *fiber.Ctx) error {
	requestURL := fmt.Sprintf("%s/%s", h.conf.Application.CinemaServiceHost, strings.TrimPrefix(c.OriginalURL(), "/auth/"))
	c.Locals("request_url", requestURL)
	h.logger.Debug().Msg("call h.Redirect")
	responseBody, err := h.Redirect(c)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(fmt.Sprintf("error sending request to cinema service: %s", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{Error: fmt.Sprintf("error sending request to film service: %s", err.Error())})
	}

	return c.Send(responseBody)
}

// GetAllCinemaHallTypes
// @Tags         Cinema
// @Summary      Получение всех типов зала кинотеатра
// @Accept       json
// @Produce      json
// @Success      200  {array}   []entities.CinemaHallType  "Список всех типов зала кинотеатра"
// @Failure      500  {object}  entities.Error  "Ошибка на стороне сервера"
// @Router       /auth/cinema/hall/types [get]
// @Security ApiKeyAuth
func (h *Handler) GetAllCinemaHallTypes(c *fiber.Ctx) error {
	requestURL := fmt.Sprintf("%s/%s", h.conf.Application.CinemaServiceHost, strings.TrimPrefix(c.OriginalURL(), "/auth/"))
	c.Locals("request_url", requestURL)
	h.logger.Debug().Msg("call h.Redirect")
	responseBody, err := h.Redirect(c)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(fmt.Sprintf("error sending request to cinema service: %s", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{Error: fmt.Sprintf("error sending request to film service: %s", err.Error())})
	}

	return c.Send(responseBody)
}

// CreateCinema
// @Tags         Cinema
// @Summary      Создание кинотеатра
// @Accept multipart/form-data
// @Produce      json
// @Param name formData string true "Название" example("Кварц")
// @Param description formData string true "Описание" example("Привет Подольск")
// @Param photo formData file false "Фото кинотеатра"
// @Param address formData string true "Адрес" example("Гевинская 9/29")
// @Param email formData string true "Контактная почта" example("meow@meow.meow")
// @Param phone formData string true "Контактный номер телефона" example("+7(952)812-02-02")
// @Param condition_id formData int true "ID состояния"
// @Param category_id formData int true "ID категории"
// @Success      200 {object} entities.ID "Кинотеатр успешно создан"
// @Failure      400 {object} entities.Error "Некорректный запрос"
// @Failure      401 {object} entities.Error "Пользователь не авторизован"
// @Failure      403 {object} entities.Error "Недостаточно прав для запроса"
// @Failure      500 {object} entities.Error "Ошибка на стороне сервера"
// @Router       /auth/cinema [post]
// @Security ApiKeyAuth
func (h *Handler) CreateCinema(c *fiber.Ctx) error {
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

	requestURL := fmt.Sprintf("%s/%s", h.conf.Application.CinemaServiceHost, strings.TrimPrefix(c.OriginalURL(), "/auth/"))
	c.Locals("request_url", requestURL)
	h.logger.Debug().Msg("call h.Redirect")
	responseBody, err := h.Redirect(c)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(fmt.Sprintf("error sending request to cinema service: %s", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{Error: fmt.Sprintf("error sending request to film service: %s", err.Error())})
	}

	return c.Send(responseBody)
}

// GetAllCinemasAddressName
// @Tags         Cinema
// @Summary      Получение всех названий и адресов кинотеатров
// @Accept       json
// @Produce      json
// @Success      200  {array}   []entities.CinemaAddressName  "Список всех названий и адресов кинотеатров"
// @Failure      500  {object}  entities.Error  "Ошибка на стороне сервера"
// @Router       /auth/cinema/address_name [get]
// @Security ApiKeyAuth
func (h *Handler) GetAllCinemasAddressName(c *fiber.Ctx) error {
	requestURL := fmt.Sprintf("%s/%s", h.conf.Application.CinemaServiceHost, strings.TrimPrefix(c.OriginalURL(), "/auth/"))
	c.Locals("request_url", requestURL)
	h.logger.Debug().Msg("call h.Redirect")
	responseBody, err := h.Redirect(c)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(fmt.Sprintf("error sending request to cinema service: %s", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{Error: fmt.Sprintf("error sending request to film service: %s", err.Error())})
	}

	return c.Send(responseBody)
}

// GetAllCinemaHallsByID
// @Tags         Cinema
// @Summary      Получение всех залов кинотеатра по его ID
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID кинотеатра"
// @Success      200  {object}  []entities.GetCinemaHall  "Залы кинотеатра успешно получены"
// @Failure      400  {object}  entities.Error  "Некорректный запрос"
// @Failure      404  {object}  entities.Error  "Зал не найден"
// @Failure      500  {object}  entities.Error  "Ошибка на стороне сервера"
// @Router       /auth/cinema/halls/{id} [get]
// @Security     ApiKeyAuth
func (h *Handler) GetAllCinemaHallsByID(c *fiber.Ctx) error {
	requestURL := fmt.Sprintf("%s/%s", h.conf.Application.CinemaServiceHost, strings.TrimPrefix(c.OriginalURL(), "/auth/"))
	c.Locals("request_url", requestURL)
	h.logger.Debug().Msg("call h.Redirect")
	responseBody, err := h.Redirect(c)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(fmt.Sprintf("error sending request to cinema service: %s", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{Error: fmt.Sprintf("error sending request to film service: %s", err.Error())})
	}

	return c.Send(responseBody)
}

// GetCinemaByID
// @Tags         Cinema
// @Summary      Получение полной информации о кинотеатре по его ID
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID кинотеатра"
// @Success      200  {object}  entities.GetCinema  "Полная информация о кинотеатре получена"
// @Failure      400  {object}  entities.Error  "Некорректный запрос"
// @Failure      404  {object}  entities.Error  "Зал не найден"
// @Failure      500  {object}  entities.Error  "Ошибка на стороне сервера"
// @Router       /auth/cinema/id/{id} [get]
// @Security     ApiKeyAuth
func (h *Handler) GetCinemaByID(c *fiber.Ctx) error {
	requestURL := fmt.Sprintf("%s/%s", h.conf.Application.CinemaServiceHost, strings.TrimPrefix(c.OriginalURL(), "/auth/"))
	c.Locals("request_url", requestURL)
	h.logger.Debug().Msg("call h.Redirect")
	responseBody, err := h.Redirect(c)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(fmt.Sprintf("error sending request to cinema service: %s", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{Error: fmt.Sprintf("error sending request to film service: %s", err.Error())})
	}

	return c.Send(responseBody)
}

// UpdateCinema
// @Tags         Cinema
// @Summary      Обновление кинотеатра
// @Accept multipart/form-data
// @Produce      json
// @Param        id                formData      int       true  "ID кинотеатра"
// @Param name formData string true "Название" example("Кварц")
// @Param description formData string true "Описание" example("Привет Подольск")
// @Param photo formData file false "Фото кинотеатра"
// @Param address formData string true "Адрес" example("Гевинская 9/29")
// @Param email formData string true "Контактная почта" example("meow@meow.meow")
// @Param phone formData string true "Контактный номер телефона" example("+7(952)812-02-02")
// @Param condition_id formData int true "ID состояния"
// @Param category_id formData int true "ID категории"
// @Success      200 {object} entities.ID "Кинотеатр успешно создан"
// @Failure      400 {object} entities.Error "Некорректный запрос"
// @Failure      401 {object} entities.Error "Пользователь не авторизован"
// @Failure      403 {object} entities.Error "Недостаточно прав для запроса"
// @Failure      500 {object} entities.Error "Ошибка на стороне сервера"
// @Router       /auth/cinema [put]
// @Security ApiKeyAuth
func (h *Handler) UpdateCinema(c *fiber.Ctx) error {
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

	requestURL := fmt.Sprintf("%s/%s", h.conf.Application.CinemaServiceHost, strings.TrimPrefix(c.OriginalURL(), "/auth/"))
	c.Locals("request_url", requestURL)
	h.logger.Debug().Msg("call h.Redirect")
	responseBody, err := h.Redirect(c)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(fmt.Sprintf("error sending request to cinema service: %s", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{Error: fmt.Sprintf("error sending request to film service: %s", err.Error())})
	}

	return c.Send(responseBody)
}

// DeleteCinema
// @Tags         Cinema
// @Summary      Удаление кинотеатра
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID кинотеатра"
// @Success      200  {object}  entities.ID  "Кинотеатр успешно удален"
// @Failure      400  {object}  entities.Error  "Некорректный запрос"
// @Failure      404  {object}  entities.Error  "Кинотеатр не найден"
// @Failure      500  {object}  entities.Error  "Ошибка на стороне сервера"
// @Router       /auth/cinema/{id} [delete]
// @Security     ApiKeyAuth
func (h *Handler) DeleteCinema(c *fiber.Ctx) error {
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

	requestURL := fmt.Sprintf("%s/%s", h.conf.Application.CinemaServiceHost, strings.TrimPrefix(c.OriginalURL(), "/auth/"))
	c.Locals("request_url", requestURL)
	h.logger.Debug().Msg("call h.Redirect")
	responseBody, err := h.Redirect(c)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(fmt.Sprintf("error sending request to cinema service: %s", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{Error: fmt.Sprintf("error sending request to film service: %s", err.Error())})
	}

	return c.Send(responseBody)
}
