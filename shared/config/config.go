package config

import (
	"reflect"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type EnvConfig struct {
	PORT            string `mapstructure:"PORT"               default:"8000"`
	DBHost          string `mapstructure:"DB_HOST"            default:"localhost"`
	DBUser          string `mapstructure:"DB_USER"            default:""`
	DBPassword      string `mapstructure:"DB_PASSWORD"        default:""`
	DBName          string `mapstructure:"DB_NAME"            default:""`
	DBPort          string `mapstructure:"DB_PORT"            default:""`
	ENV             string `mapstructure:"ENV"                default:""`
	IsPrettyLogging bool   `mapstructure:"IS_PRETTY_LOGGING"  default:"true"`
	SecretKey       string `mapstructure:"SECRET_KEY"         default:"this_is_not_secret"`
}

func NewEnvConfig(log *logrus.Logger) (*EnvConfig, error) {
	var config EnvConfig
	viper.AddConfigPath("../../")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Errorf("error while reading config: %s", err.Error())
		return nil, err
	}

	fields := reflect.VisibleFields(reflect.TypeOf(EnvConfig{}))
	for _, field := range fields {
		key := field.Tag.Get("mapstructure")
		def := field.Tag.Get("default")

		viper.SetDefault(key, def)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Errorf("error while unmarshall config: %s", err.Error())
		return nil, err
	}

	return &config, nil
}