package config

import (
	"sync"
)

type Config struct {
	Environment     string `env:"ENVIRONMENT"`
	MongoDbUrl      string `env:"MONGO_DB_URL"`
	MongoDbUsername string `env:"MONGO_DB_USERNAME"`
	MongoDbPassword string `env:"MONGO_DB_PASSWORD"`
	MongoDbDatabase string `env:"MONGO_DB_DATABASE"`
	AppName         string `env:"APP_NAME"`
	ApiMode         string `env:"API_MODE" binding:"optional"`
	ServerAddress   string `env:"SERVER_ADDRESS"`
	HttpPort        string `env:"HTTP_PORT"`
}

var once sync.Once
var instance *Config

func GetInstance() *Config {
	if instance == nil {
		once.Do(func() {
			config := setup()
			instance = &config
		})
	}

	return instance
}
