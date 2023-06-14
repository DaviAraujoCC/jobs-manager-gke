package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	SecretKey string `mapstructure:"secret_key"`
	LogLevel  string `mapstructure:"log_level"`
	HTTPPort  string `mapstructure:"http_port"`
}

func New() (*Config, error) {
	viper.SetDefault("HTTP_PORT", "8080")

	viper.SetConfigType("env")

	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Println("config: .env file not found")
	}

	viper.AutomaticEnv()

	cfg := &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil

}
