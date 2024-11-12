package pkg

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken generate a valid jwt token for 10 minutes
func GenerateToken(userID uint, secretKey []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(10 * time.Minute).Unix(), // expire in 10 minutes
	})
	return token.SignedString(secretKey)
}
