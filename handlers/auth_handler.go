package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp": time.Now().Add(time.Hour + 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	return token.SignedString(jwtSecret)
}

func Login(c echo.Context) error {
	req := new(LoginRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Request"})
	}

	// Mock data 
	if req.Email == "admin@example.com" && req.Password == "password" {
		token, err := GenerateJWT(req.Email)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
		}
		return c.JSON(http.StatusOK, map[string]string{"token": token})
	}

	return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
}