package authlib

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

func jwtKeyFunc(t *jwt.Token) (interface{}, error) {
	return JWT_KEY, nil
}

/*
r
*/
func ParseJWT(tokenString string) (JWTClaim, error) {
	claims := JWTClaim{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, jwtKeyFunc)

	if err != nil {
		v, _ := err.(*jwt.ValidationError)
		switch v.Errors {

		case jwt.ValidationErrorSignatureInvalid:
			return claims, errors.New("invalid signature (try parsing as oauth idtoken)")

		case jwt.ValidationErrorExpired:
			return claims, errors.New("token expired")

		default:
			return claims, errors.New("an error occurred")
		}
	}

	if !token.Valid {
		return claims, errors.New("token invalid")
	}

	return claims, nil
}
