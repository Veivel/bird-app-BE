package authservices

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

func jwtKeyFunc(t *jwt.Token) (interface{}, error) {
	return JWT_KEY, nil
}

func ParseJWT(tokenString string) {
	claims := JWTClaim{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, jwtKeyFunc)

	if err != nil {
		v, _ := err.(*jwt.ValidationError)
		switch v.Errors {

		case jwt.ValidationErrorSignatureInvalid:
			newToken, err := verifyIdToken(token)
			if err != nil {
				// do something to return error
				fmt.Println("")
			} else {
				fmt.Println(newToken)
			}

		case jwt.ValidationErrorExpired:
			fmt.Println("")

		default:
			fmt.Println("")
		}
	}

	if !token.Valid {
		fmt.Println("")
	}

}
