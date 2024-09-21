package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

var cfg *Config
var cfgError error

func init() {
	cfg, cfgError = NewConfig()
	if cfgError != nil {
		log.Fatalf("Config error: %s", cfgError)
	}
}

func GetConfig() *Config {
	return cfg
}

type (
	// Config -.
	Config struct {
		App
	}

	App struct {
		IP   string `env:"HTTP_SERVER_IP" env-default:"localhost"`
		Port string `env:"PORT" env-default:"8080"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./.env", cfg)
	if err != nil {
		fmt.Printf("ReadConfig error: %s, try to get from system env variables \n", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
