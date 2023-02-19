package authcontroller

import (
	"bird-app/models"
	"bird-app/services"
	"bird-app/services/authservices"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var user models.User
	var body models.UserAuth
	usersCollection := services.DB.Collection("users")

	c.BindJSON(&body)

	result := usersCollection.FindOne(context.Background(), bson.D{{"username", body.Username}})
	if result.Err() != nil {
		c.JSON(400, gin.H{
			"message": "Invalid credentials.",
		})
	} else if result.Decode(&user); bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)) != nil {
		c.JSON(400, gin.H{
			"message": "Invalid credentials.",
		})
	} else {
		var expTime time.Time
		if body.RememberMe {
			expTime = time.Now().Add(time.Hour * 24 * 7)
		} else {
			expTime = time.Now().Add(time.Hour * 2)
		}

		claims := authservices.JWTClaim{
			Username: user.Username,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "ristekbirdapp",
				ExpiresAt: jwt.NewNumericDate(expTime),
			},
		}

		tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
		token, err := tokenAlgo.SignedString(authservices.JWT_KEY)

		if err != nil {
			c.JSON(400, gin.H{
				"message": "Error generating JWT",
			})
		} else {
			c.JSON(200, gin.H{
				"token":   token,
				"message": "Login successful.",
			})
		}
	}
}

func Logout(c *gin.Context) {
	// force token to expire?
}

func Register(c *gin.Context) {
	usersCollection := services.DB.Collection("users")
	var body models.User

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
		enc, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println("Error generating a hash with bcrypt.")
			panic(err)
		}

		user := models.User{
			Username:  body.Username,
			Email:     body.Email,
			Password:  string(enc),
			CreatedAt: time.Now(),
			Avatar:    fmt.Sprintf("%s/tr:w-300,tr:h-300,BirdApp-avatars/default.jpeg", os.Getenv("IMAGEKIT_ENDPOINT_URL")),
		}

		usersCollection.InsertOne(context.Background(), user)

		c.JSON(http.StatusCreated, gin.H{
			"message": "Successfully registered user.",
			"data": map[string]string{
				"username": user.Username,
				"email":    user.Email,
			},
		})
	}
}
