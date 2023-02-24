package profilecontroller

import (
	"bird-app/lib"
	"bird-app/models"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// view user profile (public)
func View(c *gin.Context) {
	var user models.User

	username := c.Param("username")
	result := lib.DB.Collection("users").FindOne(context.Background(), bson.D{{"username", username}})
	if result.Err() != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "User not found.",
		})
	}

	result.Decode(&user)

	c.JSON(200, gin.H{
		"data": map[string]string{
			"username":   user.Username,
			"avatar":     user.Avatar,
			"created_at": user.CreatedAt.String(),
		},
	})
}

// view user posts (public)
func ViewPosts(c *gin.Context) {
	var user models.User
	var posts []models.Post

	username := c.Param("username")
	result := lib.DB.Collection("users").FindOne(context.Background(), bson.D{{"username", username}})
	if result.Err() != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "User not found.",
		})
		return
	}

	result.Decode(&user)

	cursor, err := lib.DB.Collection("posts").Find(context.Background(), bson.D{{"author", username}})
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "User has no posts.",
		})
		return
	}

	var post models.Post
	for cursor.TryNext(context.Background()) {
		cursor.Decode(&post)
		posts = append(posts, post)
		cursor.Next(context.Background())
	}

	c.JSON(200, gin.H{
		"data": posts,
	})
}
