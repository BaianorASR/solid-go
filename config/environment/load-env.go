package config

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

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
