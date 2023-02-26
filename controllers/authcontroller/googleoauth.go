package authcontroller

import (
	"bird-app/lib"
	"bird-app/lib/authlib"
	"bird-app/models"
	"bird-app/services/authservices"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"go.mongodb.org/mongo-driver/bson"
)

func GoogleInit(c *gin.Context) {
	gothic.Store = authlib.GetCookieStore(c.Request.URL.String())
	goth.UseProviders(authlib.GetGoogleProvider(os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET")))

	c.Request = c.Request.WithContext(context.WithValue(context.Background(), "provider", "google"))
	// try to get the user without re-authenticating
	if gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request); err == nil {
		c.JSON(200, gin.H{
			"user": gothUser,
		})
	} else {
		gothic.BeginAuthHandler(c.Writer, c.Request)
	}
}

func GoogleCallback(c *gin.Context) {
	var user models.User

	users := lib.DB.Collection("users")
	gothic.Store = authlib.GetCookieStore(c.Request.URL.String())
	goth.UseProviders(authlib.GetGoogleProvider(os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET")))

	gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		fmt.Println("err:", gothUser, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Could not complete user authentication.",
			"error":   err.Error(),
		})
		return
	}

	if users.FindOne(context.Background(), bson.D{{"email", gothUser.Email}}).Decode(&user); user.Email == "" {
		user := models.User{
			Username:  gothUser.Name,
			Email:     gothUser.Email,
			Password:  "None",
			CreatedAt: time.Now(),
			Avatar:    gothUser.AvatarURL,
		}

		authservices.RegisterUser(user)

		c.JSON(200, gin.H{
			"user":    user,
			"token":   gothUser.AccessToken,
			"message": "Succesfully registered user with Google.",
			"raw":     gothUser,
		})
		return
	}

	users.FindOne(
		context.Background(),
		bson.D{{"email", gothUser.Email}},
	)

	c.JSON(200, gin.H{
		"user":    user,
		"token":   gothUser.AccessToken,
		"message": "Successfully logged in with Google.",
		"raw":     gothUser,
	})
}
