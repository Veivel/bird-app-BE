package authservices

import (
	"bird-app/lib/authlib"
	"bird-app/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(credentials models.UserAuth, user models.User) (token string, expTime time.Time, err error) {
	if credentials.RememberMe {
		expTime = time.Now().Add(time.Hour * 24 * 7)
	} else {
		expTime = time.Now().Add(time.Hour * 2)
	}

	claims := authlib.JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "ristekbirdapp",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	token, err = tokenAlgo.SignedString(authlib.JWT_KEY)

	return token, expTime, err
}
