package config

import (
	"fmt"
	"os"
)

func ReadEnv(key string) (string, error) {
	value := os.Getenv(key)
	
	if value == "" {
		return "", fmt.Errorf("environment variable %s does not exist", key)
	}

	return value, nil
}