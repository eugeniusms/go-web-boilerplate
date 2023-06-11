package middleware

import (
	"fmt"
	"go-web-boilerplate/shared/config"

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
	claims, err := m.getToken(m.Env.SecretKey, c)

	if err != nil {
		return err
	}

	c.Locals("id", uint(claims["id"].(float64)))

	return c.Next()
}

func (m *Middleware) getToken(secret string, c *fiber.Ctx) (jwt.MapClaims, error) {
	header := c.Get("Authorization", "")

	if header == "" {
		return nil, fiber.ErrUnauthorized
	}

	token, err := jwt.Parse(header, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fiber.ErrUnauthorized
	}

	m.log.Infof("Attempting auth check with header claims: %s", claims)

	exp, ok := claims["exp"].(float64)
	if !ok || carbon.Now().Timestamp() > int64(exp) {
		return nil, fiber.ErrUnauthorized
	}

	return claims, nil
}

func NewMiddleware(env *config.EnvConfig, log *logrus.Logger) *Middleware {
	return &Middleware{
		Env: env,
		log: log,
	}
}