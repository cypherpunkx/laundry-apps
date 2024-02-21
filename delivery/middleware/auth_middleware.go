package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"enigmacamp.com/enigma-laundry-apps/utils/security"
	"github.com/gin-gonic/gin"
)

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var header authHeader
		if err := c.ShouldBindHeader(&header); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message" : "Unauthorized"})
			c.Abort()
			return
		}
		token := strings.Replace(header.AuthorizationHeader, "Bearer ","",1)
		if token == ""{
			c.JSON(http.StatusUnauthorized, gin.H{"message" : "Unauthorized"})
			c.Abort()
			return
		}
		fmt.Println("Token ", token)

		claims, err := security.VerifyAccessToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message" : "Unauthorized",
				"error" : err.Error(),
			})
			c.Abort()
			return
		}
		fmt.Println("Claims :",claims["username"])
		c.Next()
	}
}