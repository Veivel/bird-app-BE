package commentcontroller

import (
	"bird-app/lib"
	"bird-app/models"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// List all comments under a post
func Index(c *gin.Context) {
	postUuid := c.Param("postUuid")
	var post models.Post

	lib.DB.Collection("posts").FindOne(
		context.Background(),
		bson.D{{"uuid", postUuid}},
	).Decode(&post)

	c.JSON(200, gin.H{
		"data": post.Comments,
	})

}

// Create a comment under a post
func Create(c *gin.Context) {
	var body models.Comment
	var post models.Post

	postUuid := c.Param("postUuid")
	posts := lib.DB.Collection("posts")
	criteria := bson.D{{"uuid", postUuid}}

	// Create the Comment struct
	c.BindJSON(&body)
	username, _ := c.Get("username")
	avatar, _ := c.Get("avatar")
	body.Author = username.(string)
	body.AuthorAvatar = avatar.(string)
	body.CreatedAt = time.Now()

	// Find existing post
	result := posts.FindOne(
		context.Background(),
		criteria,
	)
	if result.Err() != nil {
		c.AbortWithStatusJSON(404, gin.H{
			"message": "Could not find post with specified UUID.",
		})
		return
	}
	result.Decode(&post)

	// update the existing post with the new comment
	var newComments []models.Comment
	if post.Comments != nil {
		newComments = append(post.Comments, body)
	} else {
		newComments = append([]models.Comment{}, body)
	}
	posts.FindOneAndUpdate(
		context.Background(),
		criteria,
		bson.M{"$set": bson.M{"comments": newComments}},
	)

	c.JSON(201, gin.H{
		"post":    postUuid,
		"message": "Comment created",
		"data":    body,
	})
}
