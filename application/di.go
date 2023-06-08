package application

import (
	"go-web-boilerplate/application/auth"
	"go-web-boilerplate/application/healthcheck"

	"github.com/pkg/errors"
	"go.uber.org/dig"
)

type Holder struct {
	dig.In
	HealthcheckService healthcheck.Service
	AuthService        auth.Service
}

func Register(container *dig.Container) error {
	if err := container.Provide(healthcheck.NewHealthcheckService); err != nil {
		return errors.Wrap(err, "failed to provide healthcheck service")
	}

	if err := container.Provide(auth.NewAuthService); err != nil {
		return errors.Wrap(err, "failed to provide auth service")
	}

	return nil
}