package authservices

import (
	"github.com/golang-jwt/jwt/v4"
)

var JWT_KEY = []byte("QWEIQWCIJXZ")

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}
