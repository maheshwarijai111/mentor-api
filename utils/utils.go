package utils

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type APIResponse struct {
	Status  int         `json:"status"`          // HTTP Status Code
	Message string      `json:"message"`         // Response message
	Data    interface{} `json:"data,omitempty"`  // Success data (optional)
	Error   interface{} `json:"error,omitempty"` // Error details (optional)
}

func RespondJSON(c *gin.Context, status int, message string, data interface{}, err interface{}) {
	response := APIResponse{
		Status:  status,
		Message: message,
		Data:    data,
		Error:   err,
	}
	c.JSON(status, response)
}

var jwtSecret = []byte("jaiprakashmaheshwari")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateJWT generates a JWT token for a user
func GenerateJWT(username string) (string, error) {
	claims := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token expires in 24 hours
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ParseJWT parses and validates a JWT token
func ParseJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
