package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

const (
	MapClientApiKey = "MAP_CLIENT_API_KEY"
)

func SetupEnv(filepath *string) error {
	var err error
	if filepath != nil {
		err = godotenv.Load(*filepath)
	} else {
		err = godotenv.Load()
	}

	if err != nil {
		return fmt.Errorf("error loading env file: %v", err)
	}
	return nil
}

func GetMapClientApiKey() string {
	return os.Getenv(MapClientApiKey)
}
