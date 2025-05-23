package handler

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"kino/internal/shared/entities"
	"kino/internal/shared/log"
	"kino/pkg/util"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
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
	h.logger.Debug().Caller().Msg("body parse")
	var f entities.CreateCinema
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

	if f.Address == "" {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg("empty field 'address'")
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "empty field 'address'"})
	}

	if f.Email == "" {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg("empty field 'email'")
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "empty field 'email'"})
	}

	if f.Phone == "" {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg("empty field 'phone'")
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "empty field 'phone'"})
	}

	if f.ConditionID <= 0 {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Int("id", f.ConditionID).Msg("invalid field 'condition_id'")
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "invalid field 'condition_id'"})
	}

	if f.CategoryID <= 0 {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Int("id", f.CategoryID).Msg("invalid field 'category_id'")
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "invalid field 'category_id'"})
	}

	file, err := c.FormFile("photo")
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
	err = h.repository.S3.FPutObject(context.Background(), "cinema-media", fileName, filePath, file.Header.Get("Content-Type"))
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

	h.logger.Debug().Msg("call h.repository.DB.CreateCinema")
	id, err := h.repository.DB.CreateCinema(&f)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(fmt.Sprintf("error creating cinema: %s", err.Error()))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("error creating cinema: %s", err.Error())})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg(fmt.Sprintf("successful creating cinema with id=%v", id))
	return c.Status(fiber.StatusOK).JSON(entities.ID{ID: id})
}

func (h *Handler) CreateCinemaHall(c *fiber.Ctx) error {
	h.logger.Debug().Caller().Msg("body parse")
	var f entities.CinemaHall
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

	if f.Capacity <= 0 {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Int("id", f.Capacity).Msg("invalid field 'capacity'")
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "invalid field 'capacity'"})
	}

	if f.TypeID <= 0 {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Int("id", f.TypeID).Msg("invalid field 'type_id'")
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "invalid field 'type_id'"})
	}

	h.logger.Debug().Msg("call h.repository.DB.CreateCinemaHall")
	id, err := h.repository.DB.CreateCinemaHall(&f)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(fmt.Sprintf("error updating cinema hall: %s", err.Error()))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("error updating cinema hall: %s", err.Error())})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg(fmt.Sprintf("successful updating cinema hall with id=%v", id))
	return c.Status(fiber.StatusOK).JSON(entities.ID{ID: id})
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
		logEvent.Msg(fmt.Sprintf("invalid cinema ID: %v", err))
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "Invalid cinema ID"})
	}

	h.logger.Debug().Msg("calling h.repository.DB.GetAllCinemaHallsByID")
	cinemaHalls, err := h.repository.DB.GetAllCinemaHallsByID(id)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(fmt.Sprintf("error getting all cinema halls by id: %s", err.Error()))
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: fmt.Sprintf("error getting all cinema halls by id: %s", err.Error())})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("successful getting all cinema halls by id")
	return c.Status(fiber.StatusOK).JSON(cinemaHalls)
}

func (h *Handler) GetCinemaHallByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(fmt.Sprintf("invalid cinema hall ID: %v", err))
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "Invalid cinema hall ID"})
	}

	h.logger.Debug().Msg("calling h.repository.DB.GetCinemaHallByID")
	cinemaHall, err := h.repository.DB.GetCinemaHallByID(id)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(fmt.Sprintf("error getting cinema hall by id: %s", err.Error()))
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: fmt.Sprintf("error getting cinema hall by id: %s", err.Error())})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("successful getting cinema hall by id")
	return c.Status(fiber.StatusOK).JSON(cinemaHall)
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

	re := regexp.MustCompile("^(https?|ftp):\\/\\/[^\\s/$.?#].[^\\s]*$")

	if !re.MatchString(cinema.Photo) && cinema.Photo != "" {
		url, err := h.repository.S3.PresignedGetObject(context.Background(), "cinema-media", cinema.Photo, 7*24*time.Hour)
		if err != nil {
			logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
			logEvent.Msg(fmt.Sprintf("error getting presigned object %s from minio: %s", cinema.Photo, err.Error()))
			return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: fmt.Sprintf("error getting presigned object %s from minio: %s", cinema.Photo, err.Error())})
		}

		cinema.Photo = url.String()
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("successful getting cinema by id")
	return c.Status(fiber.StatusOK).JSON(cinema)
}

func (h *Handler) GetAllCinemas(c *fiber.Ctx) error {
	h.logger.Debug().Msg("calling h.repository.DB.GetAllCinemas")
	cinemas, err := h.repository.DB.GetAllCinemas()
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(fmt.Sprintf("error getting all cinemas: %s", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{Error: fmt.Sprintf("error getting all cinemas: %s", err.Error())})
	}

	re := regexp.MustCompile("^(https?|ftp):\\/\\/[^\\s/$.?#].[^\\s]*$")
	for i := range cinemas {
		if !re.MatchString(cinemas[i].Photo) && cinemas[i].Photo != "" {
			url, err := h.repository.S3.PresignedGetObject(context.Background(), "cinema-media", cinemas[i].Photo, 7*24*time.Hour)
			if err != nil {
				logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
				logEvent.Msg(fmt.Sprintf("error getting presigned object %s from minio: %s", cinemas[i].Photo, err.Error()))
				return c.Status(fiber.StatusInternalServerError).JSON(entities.Error{Error: fmt.Sprintf("error getting presigned object %s from minio: %s", cinemas[i].Photo, err.Error())})
			}
			cinemas[i].Photo = url.String()
		}
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("successful getting all cinemas")
	return c.Status(fiber.StatusOK).JSON(cinemas)
}

func (h *Handler) UpdateCinema(c *fiber.Ctx) error {
	h.logger.Debug().Caller().Msg("body parse")
	var f entities.Cinema
	if err := c.BodyParser(&f); err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: err.Error()})
	}

	if f.ID < 1 {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Int("id", f.ID).Msg("invalid field 'id'")
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "invalid field 'id'"})
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

	if f.Address == "" {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg("empty field 'address'")
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "empty field 'address'"})
	}

	if f.Email == "" {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg("empty field 'email'")
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "empty field 'email'"})
	}

	if f.Phone == "" {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg("empty field 'phone'")
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "empty field 'phone'"})
	}

	if f.ConditionID <= 0 {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Int("id", f.ConditionID).Msg("invalid field 'condition_id'")
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "invalid field 'condition_id'"})
	}

	if f.CategoryID <= 0 {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Int("id", f.CategoryID).Msg("invalid field 'category_id'")
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "invalid field 'category_id'"})
	}

	cinemaDB, err := h.repository.DB.GetCinemaByID(f.ID)
	if err != nil {
		status := fiber.StatusInternalServerError
		if strings.Contains(err.Error(), "not found") {
			status = fiber.StatusNotFound
		}

		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: status})
		logEvent.Msg(fmt.Sprintf("error updating cinema: error getting cinema by id: %v", err))
		return c.Status(status).JSON(entities.Error{Error: err.Error()})
	}

	file, err := c.FormFile("photo")
	if err != nil {
		f.Photo = cinemaDB.Photo
	} else {
		re := regexp.MustCompile("^(https?|ftp):\\/\\/[^\\s/$.?#].[^\\s]*$")

		if !re.MatchString(cinemaDB.Photo) {
			h.logger.Debug().Caller().Msg("call h.repository.S3.RemoveObject")
			err = h.repository.S3.RemoveObject(context.Background(), "cinema-media", cinemaDB.Photo)
			if err != nil {
				logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
				logEvent.Msg(fmt.Sprintf("error removing file %s from minio: %s", cinemaDB.Photo, err.Error()))
				return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: fmt.Sprintf("error removing file %s from minio: %s", cinemaDB.Photo, err.Error())})
			}
		}

		filePath := fmt.Sprintf("./tmp/%s", file.Filename)

		h.logger.Debug().Caller().Msg("save file")
		c.SaveFile(file, filePath)
		fileName := util.GenerateRandomFileName(filepath.Ext(file.Filename))
		f.Photo = fileName

		h.logger.Debug().Caller().Msg("call h.repository.S3.FPutObject")
		err = h.repository.S3.FPutObject(context.Background(), "cinema-media", fileName, filePath, file.Header.Get("Content-Type"))
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
	}

	h.logger.Debug().Msg("call h.repository.DB.UpdateCinema")
	err = h.repository.DB.UpdateCinema(&f)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(fmt.Sprintf("error updating cinema: %s", err.Error()))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("error updating cinema: %s", err.Error())})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg(fmt.Sprintf("successful updating cinema with id=%v", f.ID))
	return c.Status(fiber.StatusOK).JSON(entities.ID{ID: f.ID})
}

func (h *Handler) UpdateCinemaHall(c *fiber.Ctx) error {
	h.logger.Debug().Caller().Msg("body parse")
	var f entities.CinemaHall
	if err := c.BodyParser(&f); err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: err.Error()})
	}

	if f.ID < 1 {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Int("id", f.ID).Msg("invalid field 'id'")
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "invalid field 'id'"})
	}

	if f.Name == "" {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg("empty field 'name'")
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "empty field 'name'"})
	}

	if f.Capacity <= 0 {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Int("id", f.Capacity).Msg("invalid field 'capacity'")
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "invalid field 'capacity'"})
	}

	if f.TypeID <= 0 {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Int("id", f.TypeID).Msg("invalid field 'type_id'")
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "invalid field 'type_id'"})
	}

	h.logger.Debug().Msg("call h.repository.DB.UpdateCinemaHall")
	err := h.repository.DB.UpdateCinemaHall(&f)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(fmt.Sprintf("error updating cinema hall: %s", err.Error()))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("error updating cinema hall: %s", err.Error())})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg(fmt.Sprintf("successful updating cinema hall with id=%v", f.ID))
	return c.Status(fiber.StatusOK).JSON(entities.ID{ID: f.ID})
}

func (h *Handler) DeleteCinema(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(fmt.Sprintf("invalid cinema ID: %v", err))
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "Invalid cinema ID"})
	}

	h.logger.Debug().Msg("call h.repository.DB.GetCinemaByID")
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

	h.logger.Debug().Msg("call h.repository.DB.DeleteCinema")
	if err := h.repository.DB.DeleteCinema(id); err != nil {
		status := fiber.StatusInternalServerError
		if strings.Contains(err.Error(), "not found") {
			status = fiber.StatusNotFound
		}

		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: status})
		logEvent.Msg(fmt.Sprintf("error deleting cinema: %v", err))
		return c.Status(status).JSON(entities.Error{Error: err.Error()})
	}

	return c.JSON(entities.ID{ID: id})
}

func (h *Handler) DeleteCinemaHall(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(fmt.Sprintf("invalid cinema  hall ID: %v", err))
		return c.Status(fiber.StatusBadRequest).JSON(entities.Error{Error: "Invalid cinema hall ID"})
	}

	h.logger.Debug().Msg("call h.repository.DB.DeleteCinemaHall")
	if err := h.repository.DB.DeleteCinemaHall(id); err != nil {
		status := fiber.StatusInternalServerError
		if strings.Contains(err.Error(), "not found") {
			status = fiber.StatusNotFound
		}

		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(), Url: c.OriginalURL(), Status: status})
		logEvent.Msg(fmt.Sprintf("error deleting cinema  hall: %v", err))
		return c.Status(status).JSON(entities.Error{Error: err.Error()})
	}

	return c.JSON(entities.ID{ID: id})
}
