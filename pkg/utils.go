package pkg

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv(envPaths ...string) {
	err := godotenv.Load(envPaths...)
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}
}
