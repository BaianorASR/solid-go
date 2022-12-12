package config

import (
	"log"
	"os"
	"strconv"
)

// GetEnv returns the environment variables
func GetEnv() *Environment {
	return &Environment{
		Port:             getEnvAsInt("PORT"),
		Env:              getEnvAsStr("ENV"),
		CaseSensitiveURL: getEnvAsBool("CASE_SENSITIVE_URL"),
	}
}

// panicEnvNotSet panics if the environment variable is not set
func panicEnvNotSet(key string) {
	log.Fatalf("Environment variable %s is not set or is invalid", key)
}

// getEnvAsStr returns the environment variable as a string
func getEnvAsStr(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panicEnvNotSet(key)
	}

	return value
}

// getEnvAsInt returns the environment variable as an integer
func getEnvAsInt(key string) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		panicEnvNotSet(key)
	}

	return value
}

// getEnvAsBool returns the environment variable as a boolean
func getEnvAsBool(key string) bool {
	value, err := strconv.ParseBool(os.Getenv(key))
	if err != nil {
		panicEnvNotSet(key)
	}

	return value
}
