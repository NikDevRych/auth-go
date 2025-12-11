package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(JWTKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	signedToken, err := token.SignedString([]byte(JWTKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
