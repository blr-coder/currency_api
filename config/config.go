package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port     string `env:"APP_PORT" env-default:"4444"`
	Postgres postgres
}

type postgres struct {
	Host     string `env:"DB_HOST" env-default:"127.0.0.1"`
	Port     string `env:"DB_PORT" env-default:"5438"`
	Name     string `env:"DB_NAME" env-default:"currency_api_db"`
	User     string `env:"DB_USER" env-default:"currency_api_db_user"`
	Password string `env:"DB_PASS" env-default:"currency_api_db_user_pass"`
}

func NewConfig() *Config {
	cfg := new(Config)
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		panic(fmt.Errorf("failed to init config: %v", err.Error()))
	}
	return cfg

}
