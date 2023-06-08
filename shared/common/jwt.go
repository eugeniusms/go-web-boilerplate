package common

import (
	"go-web-boilerplate/shared/dto"

	"github.com/golang-jwt/jwt"
	"github.com/golang-module/carbon"
	"github.com/pkg/errors"
)

func GenerateToken(secretKey string, user dto.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = carbon.Now().AddDay().Timestamp()

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", errors.Wrap(err, "failed to generate jwt token")
	}

	return signedToken, nil
}