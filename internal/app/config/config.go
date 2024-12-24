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
		Db
	}

	App struct {
		IP   string `env:"HTTP_SERVER_IP" env-default:"localhost"`
		Port string `env:"PORT" env-default:"8080"`
	}

	Db struct {
		DbHost     string `env:"DB_HOST" env-default:"localhost"`
		DbPort     string `env:"DB_PORT" env-default:"5433"`
		DbUser     string `env:"DB_USER" env-default:""`
		DbPassword string `env:"DB_PASSWORD" env-default:""`
		Dbname     string `env:"DB_NAME" env-default:""`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	envPaths := []string{
		".env",
		"../.env",
		"../../.env",
		"../../../.env",
	}

	var err error
	for _, path := range envPaths {
		err = cleanenv.ReadConfig(path, cfg)
		if err == nil {
			break
		}
	}

	if err != nil {
		fmt.Printf("ReadConfig error: %s, try to get from system env variables \n", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
