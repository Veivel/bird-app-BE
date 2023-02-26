package avatarservices

import (
	"bird-app/lib"
	"context"
	"fmt"
	"time"

	"github.com/imagekit-developer/imagekit-go"
	"github.com/imagekit-developer/imagekit-go/api/uploader"
	"go.mongodb.org/mongo-driver/bson"
)

func Upload(fileBase64 string, username string) (resp *uploader.UploadResponse, err error) {
	ik, err := imagekit.New()
	if err != nil {
		return resp, fmt.Errorf("could not create an imagekit instance: %s", err.Error())
	}

	var f bool = false
	var t bool = true
	resp, err = ik.Uploader.Upload(context.Background(), fileBase64, uploader.UploadParam{
		Folder:            "BirdApp-avatars",
		FileName:          fmt.Sprintf("avatar_%s.jpeg", username),
		Tags:              "avatar",
		UseUniqueFileName: &f,
		OverwriteFile:     &t,
	})
	if err != nil {
		return resp, fmt.Errorf("could not upload file: %s", err.Error())
	}

	// ik.Media.PurgeCache(context.Background(), media.PurgeCacheParam{
	// 	Url: resp.Data.Url,
	// })

	lib.DB.Collection("users").FindOneAndUpdate(
		context.Background(),
		bson.D{{Key: "username", Value: username}},
		bson.M{"$set": bson.M{"avatar": fmt.Sprintf("%s?updatedAt=%o", resp.Data.Url, time.Now().UnixNano())}},
	)

	return resp, nil
}
