package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

var Config Configuration

const configFilePath = "config.yaml"

type Configuration struct {
	DB struct {
		Host       string `yaml:"host" env:"DB_HOST" env-description:"DB host name"`
		Username   string `yaml:"username" env:"DB_USER" env-description:"DB user name"`
		Password   string `yaml:"password"  env:"DB_PASSWORD" env-description:"DB user password"`
		Name       string `yaml:"name" env:"DB_NAME" env-description:"DB name"`
		Collection string `yaml:"collection" env:"DB_NAME" env-description:"Collection name"`
	} `yaml:"database"`
}

func Initialize() {
	err := cleanenv.ReadConfig(configFilePath, &Config)
	if err != nil {
		panic(fmt.Sprintf("Failed to read configuration '%s', cannot continue - %s", configFilePath, err))
	}
}
