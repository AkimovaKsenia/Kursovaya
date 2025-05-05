package handler

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type tokenClaims struct {
	jwt.MapClaims
	UserId int `json:"user_id"`
}

func (h *Handler) WithJWTAuth(c *fiber.Ctx) error {
	header := c.Get("Authorization")

	if header == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing auth token"})
	}

	tokenString := strings.Split(header, " ")

	if len(tokenString) != 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid auth header"})
	}

	id, err := h.ParseToken(tokenString[1])
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}
	// Записываем id в контекст, чтобы в дальнейшем использовать в других функциях
	c.Locals("id", id)
	return c.Next()
}

func (h *Handler) GenerateToken(id, expirationTime int) (string, error) {
	claims := &tokenClaims{
		jwt.MapClaims{
			"ExpiresAt": time.Now().Add(10000 * time.Hour).Unix(),
			"IssuedAr":  time.Now().Unix(),
		},
		id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(h.conf.Application.SigningKey))
}

func (h *Handler) ParseToken(tokenString string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(h.conf.Application.SigningKey), nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, err
	}

	if time.Now().Unix() > int64(claims.MapClaims["ExpiresAt"].(float64)) {
		return 0, fmt.Errorf("token has expired")
	}

	return claims.UserId, nil
}
