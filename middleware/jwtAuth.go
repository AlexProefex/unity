package middleware

import (
	"net/http"
	"strings"
	"unity/utils"

	"github.com/gin-gonic/gin"
)

//var secretKey = "Estaesunaclavesecreta"

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the request header
		authHeader := c.GetHeader("Authorization")
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		// Verify the token
		userID, err := utils.ExtractTokenID(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Set the user ID in the context for further use in the handler
		c.Set("userID", userID)
		c.Next()
	}
}

func JwtQRAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the request header
		authHeader := c.GetHeader("Authorization")
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		// Verify the token
		userID, cantidad, puntos, producto, err := utils.ExtractQRTokenID(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Set the user ID in the context for further use in the handler
		c.Set("userID", userID)
		c.Set("cantidad", cantidad)
		c.Set("puntos", puntos)
		c.Set("producto", producto)

		c.Next()
	}
}
