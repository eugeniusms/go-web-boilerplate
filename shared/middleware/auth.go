package middleware

import (
	"fmt"
	"go-web-boilerplate/shared/config"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/golang-module/carbon"
	"github.com/sirupsen/logrus"
)

type Middleware struct {
	Env *config.EnvConfig
	log *logrus.Logger
}

func (m *Middleware) AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return fiber.ErrUnauthorized
	}

	tokenString := strings.Split(authHeader, " ")[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(m.Env.SecretKey), nil
	})

	if err != nil {
		return fiber.ErrUnauthorized
	}

	if !token.Valid {
		return fiber.ErrUnauthorized
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return fiber.ErrUnauthorized
	}

	m.log.Infof("Attempting auth check with header claims: %s", claims)

	exp, ok := claims["exp"].(float64)
	if !ok || carbon.Now().Timestamp() > int64(exp) {
		return fiber.ErrUnauthorized
	}

	c.Locals("email", claims["email"])
	c.Locals("id", claims["id"])

	return c.Next()
}

func NewMiddleware(env *config.EnvConfig, log *logrus.Logger) *Middleware {
	return &Middleware{
		Env: env,
		log: log,
	}
}