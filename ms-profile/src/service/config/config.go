package config

import (
	"sync"
)

type Config struct {
	MongoDbUrl      string `env:"MONGO_DB_URL"`
	MongoDbUsername string `env:"MONGO_DB_USERNAME"`
	MongoDbPassword string `env:"MONGO_DB_PASSWORD"`
	MongoDbDatabase string `env:"MONGO_DB_DATABASE"`
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
