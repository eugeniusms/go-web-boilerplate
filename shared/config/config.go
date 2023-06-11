package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type EnvConfig struct {
	HOST                   string `envconfig:"HOST"               default:"localhost:8000"`
	PORT                   string `envconfig:"PORT"               default:"8000"`
	DBHost                 string `envconfig:"DB_HOST"            default:"localhost"`
	DBUser                 string `envconfig:"DB_USER"            default:""`
	DBPassword             string `envconfig:"DB_PASSWORD"        default:""`
	DBName                 string `envconfig:"DB_NAME"            default:""`
	DBPort                 string `envconfig:"DB_PORT"            default:""`
	ENV                    string `envconfig:"ENV"                default:""`
	IsPrettyLogging        bool   `envconfig:"IS_PRETTY_LOGGING"  default:"true"`
	SecretKey              string `envconfig:"SECRET_KEY"         default:"this_is_not_secret"`
	SendinBlueApiKey       string `envconfig:"SENDINBLUE_API_KEY" default:""`
	ClientPasswordResetUrl string `envconfig:"CLIENT_PASSWORD_RESET_URL" default:""`
}

func NewEnvConfig(log *logrus.Logger) (*EnvConfig, error) {
	var config EnvConfig
	err := envconfig.Process("web-go-boilerplate", &config)
	if err != nil {
		log.Errorf("error while reading config: %s", err.Error())
		return nil, err
	}

	log.Infoln("importing environment variable success")

	return &config, nil
}