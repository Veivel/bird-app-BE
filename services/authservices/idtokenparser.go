package authservices

import (
	"context"
	"os"

	"github.com/golang-jwt/jwt"
	"google.golang.org/api/idtoken"
)

func VerifyIdToken(tokenString string) (*idtoken.Payload, error) {
	claims := JWTClaim{}
	token, _ := jwt.ParseWithClaims(tokenString, &claims, jwtKeyFunc)

	return idtoken.Validate(
		context.Background(),
		token.Raw,                     // raw token string
		os.Getenv("GOOGLE_CLIENT_ID"), // audience of token
	)
}
