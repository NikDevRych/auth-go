package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

var jwtTokenKey = "some-jwt-token-key-only-for-test"

func CreateToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodES256)
	signedToken, err := token.SignedString(jwtTokenKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
