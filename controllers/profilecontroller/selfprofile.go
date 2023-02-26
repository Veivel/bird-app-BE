package profilecontroller

import (
	"bird-app/lib"
	"bird-app/lib/encoder"
	"bird-app/models"
	"bird-app/services/avatarservices"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func ViewSelf(c *gin.Context) {
	var user models.User
	username, _ := c.Get("username")

	lib.DB.Collection("users").FindOne(
		context.Background(),
		bson.D{{"username", username}},
	).Decode(&user)

	user.Password = ""

	c.JSON(200, gin.H{
		"data": user,
	})
}

func ViewPostsSelf(c *gin.Context) {}

/*
edit user profile (only self)

read up:
- https://freshman.tech/snippets/go/image-to-base64/
- https://docs.imagekit.io/api-reference/upload-file-api/server-side-file-upload
*/
func EditAvatar(c *gin.Context) {
	var user models.User
	username, _ := c.Get("username")

	lib.DB.Collection("users").FindOne(
		context.Background(),
		bson.D{{"username", username}},
	).Decode(&user)

	form, err := c.MultipartForm()
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "Multipart form invalid.",
		})
		return
	}
	files := form.File["avatar"]

	for _, file := range files {
		actualFile, _ := file.Open()
		fileBase64 := encoder.GetBase64(actualFile)

		resp, err := avatarservices.Upload(fileBase64, username.(string))
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"avatar": resp.Data.Url,
		})
		return
	}
}
