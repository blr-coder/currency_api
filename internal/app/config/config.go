package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Port             string   `yaml:"port"`
	Postgres         postgres `yaml:"postgres"`
	PostgresConnLink string   `yaml:"postgres_conn_link"`
	AbstractApiKey   string   `yaml:"abstract_api_key"`
}

type postgres struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func NewConfig(configPath string) (*Config, error) {
	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err = yaml.Unmarshal(configFile, &config); err != nil {
		return nil, err
	}

	return &config, nil

}
