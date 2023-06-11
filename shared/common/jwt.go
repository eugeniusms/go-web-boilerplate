package common

import (
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

func GenerateToken(secretKey string, claims jwt.MapClaims) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", errors.Wrap(err, "failed to generate jwt token")
	}

	return signedToken, nil
}