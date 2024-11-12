package pkg

import (
	"github.com/gin-gonic/gin"
)

// AuthAuthorization is a middleware for authorization using JWT
func AuthAuthorization(secretKey []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"message": "No token provided"})
			return
		}
		userID, err := VerifyToken(token, secretKey)
		if err != nil {
			c.AbortWithStatusJSON(403, gin.H{"message": err.Error()})
			return
		}

		// Attach the userID to the request
		c.Set("userID", userID)
		c.Next()
	}
}
