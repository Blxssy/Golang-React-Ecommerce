package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Database
}

type Database struct {
	Dialect   string `yaml:"dialect" default:"postgres"`
	Host      string `yaml:"host" default:"localhost"`
	Port      string `yaml:"port"`
	Name      string `yaml:"name"`
	Username  string `yaml:"username"`
	Password  string
	Migration bool `yaml:"migration"`
}

func NewConfig() *Config {
	configPath := os.Getenv("ConfigPath")
	if configPath == "" {
		configPath = "../../configs/config.yaml"
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file not found")
	}
	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	cfg.Password = os.Getenv("DB_PASSWORD")

	return &cfg
}
