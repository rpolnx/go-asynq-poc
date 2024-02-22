package configs

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type AppProfile string

const (
	LocalProfile AppProfile = "LOCAL"
	DevProfile   AppProfile = "DEV"
	TestProfile  AppProfile = "TEST"
	ProdProfile  AppProfile = "PROD"
)

type AppConfig struct {
	AppName string      `env:"APP_NAME"`
	Host    string      `env:"HOST"`
	Port    int         `env:"PORT"`
	Profile AppProfile  `env:"PROFILE"`
	Redis   RedisConfig `env:"REDIS"`
}

type RedisConfig struct {
	Host     string `env:"REDIS_HOST"`
	Port     int    `env:"REDIS_PORT"`
	User     string `env:"REDIS_USER"`
	Password string `env:"REDIS_PASSWORD"`
	Db       int    `env:"REDIS_DB"`
}

var GlobalAppConfig *AppConfig = &AppConfig{}

func InitEnvVariables() (*AppConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	err = env.Parse(GlobalAppConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to parse environment variables: %v", err)
	}

	return GlobalAppConfig, nil
}
