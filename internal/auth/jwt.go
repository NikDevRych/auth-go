package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

var jwtTokenKey = "some-jwt-token-key-only-for-test"

func CreateToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	signedToken, err := token.SignedString([]byte(jwtTokenKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
