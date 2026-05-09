package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(filename string) error {
	log.Printf("Loading env %s...", filename)

	if err := godotenv.Load(filename); err != nil {
		return fmt.Errorf("There was an error in loading %s:\n\t%s", filename, err.Error())
	}

	log.Printf("Successfully loaded %s!", filename)
	return nil
}

func GetEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("%s is either empty or undefined in the env file", key)
	}

	return value, nil
}
