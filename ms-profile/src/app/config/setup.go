package config

import (
	"log"
	"os"
	"path/filepath"
	"reflect"

	"github.com/joho/godotenv"
	"github.com/my-storage/ms-profile/src/shared/utils"
)

func setup() Config {
	currentPath, err := utils.GetCurrentPath()
	if err != nil {
		log.Fatalf("Error on get current path: %v", err.Error())
	}

	envPath := filepath.Join(*currentPath, ".env")

	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Error loading .env file: %v", err.Error())
	}

	config := Config{}

	val := reflect.ValueOf(&config).Elem()

	for i := 0; i < val.NumField(); i++ {
		envKey := val.Type().Field(i).Tag.Get("env")
		fieldName := val.Type().Field(i).Name

		if envKey == "" {
			log.Fatalf("Error on read value from env, struct '%v' field have no 'env' tag", fieldName)
		}

		envValue := os.Getenv(envKey)

		if envValue == "" {
			bindingValue := val.Type().Field(i).Tag.Get("binding")

			if bindingValue != "optional" {
				log.Fatalf("Error on read value from env, 'env' tag '%v' in field '%v' on Config struct, not exist on env file", envKey, fieldName)
			}
		}

		val.FieldByName(fieldName).SetString(envValue)
	}

	return config
}
