package pkg_test

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/stretchr/testify/assert"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken("testuser")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestTokenValidity(t *testing.T) {
	token, err := GenerateToken("testuser")
	assert.NoError(t, err)

	claims := &Claims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	assert.NoError(t, err)
	assert.True(t, parsedToken.Valid)
	assert.Equal(t, "testuser", claims.Username)
}

func TestExpiredToken(t *testing.T) {
	expiredTime := time.Now().Add(-5 * time.Minute)
	claims := &Claims{
		Username: "testuser",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	assert.NoError(t, err)

	parsedToken, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	assert.Error(t, err)
	assert.False(t, parsedToken.Valid)
}
