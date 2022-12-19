package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type (
	// Config -.
	Config struct {
		Log     `yaml:"logger"`
		Postgre `yaml:"postgre"`
	}

	// Log -.
	Log struct {
		Level string `yaml:"log_level" env:"LOG_LEVEL"`
	}

	// Mongo
	Postgre struct {
		URL string `yaml:"url" env:"PG_URL"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	viper.AddConfigPath("../../config/")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
