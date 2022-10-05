package config

import (
	"gopkg.in/yaml.v3"
	"io/fs"
	"os"
)

type Config struct {
	Port             string `yaml:"port"`
	PostgresConnLink string `yaml:"postgres_conn_link"`
	AbstractApiKey   string `yaml:"abstract_api_key"`
}

func NewConfig(configPath string) (*Config, error) {
	var config Config

	configFile, err := os.OpenFile(configPath, os.O_RDONLY, fs.ModePerm)
	if err != nil {
		return nil, err
	}

	return &config, yaml.NewDecoder(configFile).Decode(&config)
}
