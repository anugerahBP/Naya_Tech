package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Auth(c *fiber.Ctx) error {
	header := c.Get("Authorization")
	if header == "" {
		return c.Status(401).SendString("invalid token")
	}

	headerS := strings.Split(header, " ")
	if len(headerS) != 2 {
		return c.Status(401).SendString("invalid token")
	}

	if headerS[0] != "Bearer" {
		return c.Status(401).SendString("invalid token")
	}

	token := headerS[1]
	jwk, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})
	if err != nil {
		return c.Status(401).SendString("invalid token")
	}

	c.Locals("email", jwk.Claims.(jwt.MapClaims)["email"])

	return c.Next()
}
