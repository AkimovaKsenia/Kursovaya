package log

import (
	"fmt"
	"kino/internal/shared/config"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/rs/zerolog"
)

type LogsField struct {
	Level  string
	Method string
	Url    string
	Status int
}

func InitLogger(envConf *config.Config) *zerolog.Logger {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	return &logger
}

func CreateLog(log *zerolog.Logger, field LogsField) *zerolog.Event {
	var event *zerolog.Event
	if field.Level == "Info" {
		event = log.Info()
	} else if field.Level == "Error" {
		event = log.Error()
	} else if field.Level == "Warn" {
		event = log.Warn()
	} else if field.Level == "Debug" {
		event = log.Debug()
	} else if field.Level == "Fatal" {
		event = log.Fatal()
	} else {
		fmt.Println("Unknown log level")
		return nil
	}

	event.Str("method", field.Method).Str("url", field.Url).Int("status", field.Status)

	return event
}

func RequestLogger(log *zerolog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Info().
			Str("method", c.Method()).
			Str("url", c.OriginalURL()).
			Msg("incoming request")

		start := time.Now()
		defer func() {
			if time.Since(start) > time.Second*2 {
				log.Warn().
					Str("method", c.Method()).
					Str("url", c.OriginalURL()).
					Dur("elapsed_ms", time.Since(start)).
					Msg("long response time")
			}
		}()

		return c.Next()
	}
}
