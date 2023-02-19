package authservices

import (
	"context"
	"os"

	"github.com/golang-jwt/jwt"
	"google.golang.org/api/idtoken"
)

func verifyIdToken(token *jwt.Token) (*idtoken.Payload, error) {
	return idtoken.Validate(
		context.Background(),
		token.Raw,                     // raw token string
		os.Getenv("GOOGLE_CLIENT_ID"), // audience of token
	)
}
