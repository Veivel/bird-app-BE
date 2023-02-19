package middlewares

import (
	"bird-app/models"
	"bird-app/services"
	"bird-app/services/authservices"
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func JWT(c *gin.Context) {

	// check for auth header
	authHeader := c.GetHeader("Authorization")
	if !strings.Contains(authHeader, " ") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid auth header.",
		})
		return
	}

	// check bearer token format
	_, tokenString := strings.Split(authHeader, " ")[0], strings.Split(authHeader, " ")[1]
	if tokenString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid auth header.",
		})
		return
	}

	var user models.User
	var criteria bson.D
	// parse token as jwt
	claims, err := authservices.ParseJWT(tokenString)
	if err != nil {
		// parse IDtoken as jwt (oauth2)
		idtoken, err := authservices.VerifyIdToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}
		criteria = bson.D{{"username", idtoken.Claims["name"]}}
	} else {
		criteria = bson.D{{"username", claims.Username}}
	}

	result := services.DB.Collection("users").FindOne(context.Background(), criteria)
	if result.Err() != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": result.Err().Error(),
		})
	}

	// unmarshal user object
	result.Decode(&user)
	c.Set("username", user.Username)
	c.Set("email", user.Email)
	c.Set("avatar", user.Avatar)
	c.Set("closefriends", user.CloseFriends)

	c.Next()

}
