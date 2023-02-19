package profilecontroller

import (
	"bird-app/models"
	"bird-app/services"
	"bird-app/services/encoderservice"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/imagekit-developer/imagekit-go"
	"github.com/imagekit-developer/imagekit-go/api/uploader"
	"go.mongodb.org/mongo-driver/bson"
)

func ViewSelf(c *gin.Context) {}

func ViewPostsSelf(c *gin.Context) {}

/*
edit user profile (only self)

read up:
- https://freshman.tech/snippets/go/image-to-base64/
- https://docs.imagekit.io/api-reference/upload-file-api/server-side-file-upload
*/
func EditAvatar(c *gin.Context) {
	// var body models.User
	var user models.User
	username, _ := c.Get("username")
	// avatar, _ := c.Get("avatar")

	services.DB.Collection("users").FindOne(
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
		fileBase64 := encoderservice.GetBase64(actualFile)

		ik, err := imagekit.New()
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"message": "Could not create an ImageKit instance",
				"err":     err.Error(),
			})
			return
		}

		var f bool = false
		var t bool = true
		resp, err := ik.Uploader.Upload(context.Background(), fileBase64, uploader.UploadParam{
			Folder:            "BirdApp-avatars",
			FileName:          fmt.Sprintf("avatar_%s.jpeg", username),
			Tags:              "avatar",
			UseUniqueFileName: &f,
			OverwriteFile:     &t,
		})
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"message": "Could not upload file",
				"err":     err.Error(),
			})
			return
		}

		services.DB.Collection("users").FindOneAndUpdate(
			context.Background(),
			bson.D{{"username", username}},
			bson.M{
				"$set": bson.M{"avatar": resp.Data.Url},
			},
		)

		c.JSON(200, gin.H{
			"avatar": resp.Data.Url,
		})
		return
	}
}
