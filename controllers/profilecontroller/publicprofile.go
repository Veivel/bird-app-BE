package profilecontroller

import (
	"bird-app/models"
	"bird-app/services"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// view user profile (public)
func View(c *gin.Context) {
	var user models.User

	username := c.Param("username")
	services.DB.Collection("users").FindOne(context.Background(), bson.D{{"username", username}}).Decode(&user)

	c.JSON(200, gin.H{
		"data": map[string]string{
			"username":   user.Username,
			"avatar":     user.Avatar,
			"created_at": user.CreatedAt.String(),
		},
	})
}

// view user posts (public)
func ViewPosts(c *gin.Context) {}
