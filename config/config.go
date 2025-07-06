package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type (
	GlobalErrorHandlerResp struct {
		Message string `json:"message"`
	}
)

func Config(key string) string {
	// load .env file
	err := godotenv.Load(dir(".env"))
	if err != nil {
		fmt.Printf("Error loading .env file error is %s", err.Error())
	}
	return os.Getenv(key)
}

func dir(envFile string) string {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	for {
		goModPath := filepath.Join(currentDir, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			break
		}

		parent := filepath.Dir(currentDir)
		if parent == currentDir {
			currentDir, _ = os.Getwd()
			return filepath.Join(currentDir, "go.mod")
		}
		currentDir = parent
	}

	return filepath.Join(currentDir, envFile)
}
