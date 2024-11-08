package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// AuthAuthorization is a middleware for authorization using JWT
func AuthAuthorization() gin.HandlerFunc {
	err := godotenv.Load()
	if err != nil {
		return func(c *gin.Context) {
			c.AbortWithStatusJSON(500, gin.H{"message": "Error loading .env file"})
		}
	}
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"message": "No token provided"})
			return
		}
		userID, err := VerifyToken(token)
		if err != nil {
			c.AbortWithStatusJSON(403, gin.H{"message": err.Error()})
			return
		}

		// Attach the userID to the request
		c.Set("userID", userID)
		c.Next()
	}

}
