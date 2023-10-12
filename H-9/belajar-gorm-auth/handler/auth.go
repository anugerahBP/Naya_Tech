package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RouteAuth(server fiber.Router) {
	router := server.Group("/auth")
	router.Post("login", PostLogin)
}

func PostLogin(c *fiber.Ctx) error {
	claims := jwt.MapClaims{
		"email": "alex@mail.com",
		"exp":   time.Now().Add(time.Hour * 3).Unix(),
	}
	jwk := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := jwk.SignedString([]byte(""))
	return c.JSON(fiber.Map{
		"token": token,
	})
}
