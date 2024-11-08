package pkg

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken generate a valid jwt token for 5 minutes
func GenerateToken(userID uint) (string, error) {
	secretKey := []byte(os.Getenv("JWT_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(5 * time.Minute).Unix(), // expire in 5 minutes
	})
	return token.SignedString(secretKey)
}
