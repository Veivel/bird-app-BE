package postcontroller

import (
	"bird-app/models"
	"bird-app/services"
	"context"
	"math"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Show user's "home timeline"
func Index(c *gin.Context) {
	const POSTS_PER_PAGE = 10

	var posts [POSTS_PER_PAGE]models.Post
	pageStr, _ := strconv.Atoi(c.Query("page"))
	newNum := POSTS_PER_PAGE * (int64(math.Max(1.0, float64(pageStr))) - 1)

	cursor, err := services.DB.Collection("posts").Find(context.Background(), bson.D{}, &options.FindOptions{
		Skip: &newNum,
	})
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		for i := 0; i < POSTS_PER_PAGE && cursor.TryNext(context.Background()); i++ {
			cursor.Decode(&posts[i])
		}
		c.JSON(200, gin.H{
			"data": posts,
		})
	}

}

// Submit new post
func Create(c *gin.Context) {
	// var post models.Post
	var body models.Post
	username, _ := c.Get("username")
	avatar, _ := c.Get("avatar")

	err := c.BindJSON(&body) // text, is_close_friends
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	post := models.Post{
		Uuid:           uuid.New().String(),
		Text:           body.Text,
		IsCloseFriends: body.IsCloseFriends,
		Author:         username.(string),
		AuthorAvatar:   avatar.(string),
		CreatedAt:      time.Now(),
		Comments:       []models.Comment{},
		Likes:          0,
	}

	services.DB.Collection("posts").InsertOne(context.Background(), post)

	c.JSON(201, gin.H{
		"message": "Successfully created post.",
		"data":    post,
	})
}

// Edit post content
func Edit(c *gin.Context) {

}

// Delete post
func Delete(c *gin.Context) {}
