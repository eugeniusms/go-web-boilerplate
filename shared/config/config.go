package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type EnvConfig struct {
	PORT       string `mapstructure:"PORT"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBPort     string `mapstructure:"DB_PORT"`
	ENV        string `mapstructure:"ENV"`
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

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Errorf("error while unmarshall config: %s", err.Error())
		return nil, err
	}

	return &config, nil
}