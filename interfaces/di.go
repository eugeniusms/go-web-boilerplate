package interfaces

import (
	"go-web-boilerplate/interfaces/auth"
	"go-web-boilerplate/interfaces/healthcheck"

	"github.com/pkg/errors"
	"go.uber.org/dig"
)

type Holder struct {
	dig.In
	HealthcheckViewService  healthcheck.ViewService
	AuthViewService         auth.ViewService
}

func Register(container *dig.Container) error {
	if err := container.Provide(healthcheck.NewViewService); err != nil {
		return errors.Wrap(err, "failed to provide healthcheck view service")
	}

	if err := container.Provide(auth.NewViewService); err != nil {
		return errors.Wrap(err, "failed to provide auth view service")
	}

	return nil
}