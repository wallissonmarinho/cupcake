package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
}

func Load() *Config {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatalf("erro ao ler arquivo de config: %v", err)
		}
	}

	return &Config{}
}
