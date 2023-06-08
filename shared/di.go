package shared

import (
	"go-web-boilerplate/shared/config"
	depedencies "go-web-boilerplate/shared/dependencies"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

type Holder struct {
	dig.In
	Logger *logrus.Logger
	Env    *config.EnvConfig
	Http   *fiber.App
	DB     *gorm.DB
}

func Register(container *dig.Container) error {
	if err := container.Provide(depedencies.NewLogger); err != nil {
		return errors.Wrap(err, "failed to register logger")
	}

	if err := container.Provide(config.NewEnvConfig); err != nil {
		return errors.Wrap(err, "failed to provide config")
	}

	if err := container.Provide(depedencies.NewHttp); err != nil {
		return errors.Wrap(err, "failed to provide http")
	}

	if err := container.Provide(depedencies.NewDatabase); err != nil {
		return errors.Wrap(err, "failed to provide db")
	}

	return nil
}