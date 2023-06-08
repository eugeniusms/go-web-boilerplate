package middleware

import (
	"fmt"
	"go-web-boilerplate/shared/config"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type Middleware struct {
	Env *config.EnvConfig
}

func (m *Middleware) AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return fiber.ErrUnauthorized
	}
	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(m.Env.SecretKey), nil
	})
	if err != nil {
		return fiber.ErrUnauthorized
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return fiber.ErrUnauthorized
	}

	id, ok := claims["id"].(string)
	if !ok {
		return fiber.ErrUnauthorized
	}

	email, ok := claims["email"].(string)
	if !ok {
		return fiber.ErrUnauthorized
	}

	exp, ok := claims["exp"].(int64)
	if !ok || exp < time.Now().Unix() {
		return fiber.ErrUnauthorized
	}

	c.Locals("email", email)
	c.Locals("id", id)

	return c.Next()
}

func NewMiddleware(env *config.EnvConfig) *Middleware {
	return &Middleware{
		Env: env,
	}
}