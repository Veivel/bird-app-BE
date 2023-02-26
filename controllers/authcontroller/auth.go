package authcontroller

import (
	"bird-app/lib"
	"bird-app/models"
	"bird-app/services/authservices"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var user models.User
	var body models.UserAuth
	usersCollection := lib.DB.Collection("users")

	c.BindJSON(&body)

	result := usersCollection.FindOne(context.Background(), bson.D{{"username", body.Username}})
	if result.Err() != nil {
		c.JSON(400, gin.H{
			"message": "Invalid credentials. (1)",
		})
		return
	} else if result.Decode(&user); bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)) != nil {
		c.JSON(400, gin.H{
			"message": "Invalid credentials. (2)",
		})
		return
	} else {
		token, expTime, err := authservices.GenerateToken(body, user)

		if err != nil {
			c.JSON(400, gin.H{
				"message": "Error generating JWT",
			})
			return
		} else {
			c.JSON(200, gin.H{
				"token":           token,
				"message":         "Login successful.",
				"expiration_time": expTime,
			})
			return
		}
	}
}

func Logout(c *gin.Context) {
	// force token to expire?
}

func Register(c *gin.Context) {
	usersCollection := lib.DB.Collection("users")
	var body models.UserAuth

	err := c.BindJSON(&body)
	if err != nil || body.Email == "" || body.Password == "" || body.Username == "" {
		c.JSON(400, gin.H{
			"message": "Invalid body.",
		})
	} else if user := usersCollection.FindOne(context.Background(), bson.D{{"email", body.Email}}); user.Err() == nil {
		c.JSON(400, gin.H{
			"message": "E-mail is already in use.",
		})
	} else if user := usersCollection.FindOne(context.Background(), bson.D{{"username", body.Username}}); user.Err() == nil {
		c.JSON(400, gin.H{
			"message": "Username is already in use.",
		})
	} else {
		user, _ := authservices.RegisterWithCredentials(body)

		c.JSON(http.StatusCreated, gin.H{
			"message": "Successfully registered user.",
			"data": map[string]string{
				"username": user.Username,
				"email":    user.Email,
			},
		})
	}
}
