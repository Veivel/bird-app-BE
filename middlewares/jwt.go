package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWT(c *gin.Context) {

	authHeader := c.GetHeader("Authorization")
	if !strings.Contains(authHeader, " ") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid auth header.",
		})
		return
	}

	_, tokenString := strings.Split(authHeader, " ")[0], strings.Split(authHeader, " ")[1]
	// fmt.Println("Auth with", authType, "token:", tokenString)

	if tokenString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid auth header.",
		})
		return
	}

}
