package authlib

import (
	"context"
	"os"

	"google.golang.org/api/idtoken"
)

func VerifyIdToken(tokenString string) (*idtoken.Payload, error) {
	return idtoken.Validate(
		context.Background(),
		tokenString,                   // raw token string
		os.Getenv("GOOGLE_CLIENT_ID"), // audience of token
	)
}
