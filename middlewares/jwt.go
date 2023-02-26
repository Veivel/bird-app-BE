package middlewares

import (
	"bird-app/lib"
	"bird-app/lib/authlib"
	"bird-app/models"
	"context"
	"fmt"
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
	claims, err := authlib.ParseJWT(tokenString)
	if err != nil {
		fmt.Println(tokenString)
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "No token supplied",
			})
		}

		// parse IDtoken as jwt (oauth2)
		idtoken, err := authlib.VerifyIdToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Could not verify IDToken",
				"error":   err.Error(),
			})
			return
		}
		criteria = bson.D{{Key: "username", Value: idtoken.Claims["name"]}}
	} else {
		criteria = bson.D{{Key: "username", Value: claims.Username}}
	}

	result := lib.DB.Collection("users").FindOne(context.Background(), criteria)
	if result.Err() != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Could not find user",
			"error":   result.Err().Error(),
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
