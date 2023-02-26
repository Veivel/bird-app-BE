package authservices

import (
	"bird-app/lib"
	"bird-app/models"
	"bird-app/services/avatarservices"
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func RegisterWithCredentials(credentials models.UserAuth) (user models.User, err error) {
	enc, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}

	resp, err := avatarservices.Upload(avatarservices.GetDefaultAvatar(), credentials.Username)
	if err != nil {
		return user, err
	}

	user = models.User{
		Username:  credentials.Username,
		Email:     credentials.Email,
		Password:  string(enc),
		CreatedAt: time.Now(),
		Avatar:    resp.Data.Url,
	}

	err = RegisterUser(user)

	if err != nil {
		return user, err
	} else {
		return user, nil
	}
}

func RegisterUser(user models.User) (err error) {
	_, err = lib.DB.Collection("users").InsertOne(context.Background(), user)

	return
}
