package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	CDNPort string `mapstructure:"CDNPort"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		logrus.Error("Couldn't parse .env file")
	}

	err = viper.Unmarshal(&config)
	logrus.Info("Successfully parse .env file")
	return
}
