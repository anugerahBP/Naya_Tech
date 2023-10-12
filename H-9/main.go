package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// Kunci rahasia JWT, biasanya harus disimpan dengan aman
var jwtKey = []byte("rahasia_jwt")

// Struktur untuk menyimpan informasi pengguna
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Fungsi untuk membuat token JWT
func createToken(username string) (string, error) {
	// Menyusun payload token
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token berlaku selama 24 jam
	}

	// Membuat token dengan menggunakan kunci rahasia
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Middleware untuk memeriksa token JWT dari permintaan
func authenticate(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "Token JWT tidak ditemukan"})
	}

	// Memverifikasi token JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Memeriksa metode penandatanganan
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Metode penandatanganan tidak valid")
		}
		return jwtKey, nil
	})

	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "Token JWT tidak valid"})
	}

	// Menyimpan informasi pengguna ke konteks
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "Token JWT tidak valid"})
	}
	c.Locals("username", claims["username"])

	return c.Next()
}

func main() {
	app := fiber.New()

	// Endpoint untuk login dan mendapatkan token JWT
	app.Post("/login", func(c *fiber.Ctx) error {
		var user User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Permintaan tidak valid"})
		}

		// Validasi informasi pengguna (biasanya dilakukan dengan database)
		if user.Username == "pengguna" && user.Password == "password" {
			token, err := createToken(user.Username)
			if err != nil {
				return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Gagal membuat token JWT"})
			}

			return c.Status(http.StatusOK).JSON(fiber.Map{"token": token})
		}

		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "Login gagal"})
	})

	// Endpoint yang memerlukan autentikasi dengan JWT
	app.Get("/profile", authenticate, func(c *fiber.Ctx) error {
		username := c.Locals("username").(string)
		return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Halo, " + username + "!"})
	})

	port := ":3000"
	log.Printf("Server berjalan pada port %s\n", port)
	log.Fatal(app.Listen(port))
}
