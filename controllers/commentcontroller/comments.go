package commentcontroller

import (
	"bird-app/models"
	"bird-app/services"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// List all comments under a post
func Index(c *gin.Context) {
	postUuid := c.Param("postUuid")
	var post models.Post

	services.DB.Collection("posts").FindOne(
		context.Background(),
		bson.D{{"uuid", postUuid}},
	).Decode(&post)

	c.JSON(200, gin.H{
		"data": post.Comments,
	})

}

// Create a comment under a post
func Create(c *gin.Context) {}
