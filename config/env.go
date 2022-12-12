package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
	"strconv"
)

// Environment is a struct that contains all the environment variables
type Environment struct {
	Port int `env:"PORT"`
}

type EnvironmentSource interface {
	GetEnv() *Environment
}

// projectDirName is the name of the project directory
const projectDirName = "solid-go"

// LoadEnv loads the environment variables
func LoadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

// GetEnv returns the environment variables
func GetEnv() *Environment {
	return &Environment{
		Port: getEnvAsInt("PORT"),
	}
}

// getEnvAsInt returns the environment variable as an integer
func getEnvAsInt(key string) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		log.Fatal(err)
	}

	return value
}
