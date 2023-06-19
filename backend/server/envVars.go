package server

import (
	"fmt"
	"os"
)

const defaultPort = "8080"

type envVars struct {
	Port           string
	DbUsername     string
	DbPassword     string
	DbHost         string
	DbName         string
	ClerkSecretKey string
}

func GetEnvVars() (envVars, error) {
	envVars := envVars{}
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = defaultPort
	}
	envVars.Port = port
	username, exists := os.LookupEnv("DB_USERNAME")
	if !exists {
		return envVars, fmt.Errorf("DB_USERNAME not set")
	}
	envVars.DbUsername = username
	password, exists := os.LookupEnv("DB_PASSWORD")
	if !exists {
		return envVars, fmt.Errorf("DB_PASSWORD not set")
	}
	envVars.DbPassword = password
	host, exists := os.LookupEnv("DB_HOST")
	if !exists {
		return envVars, fmt.Errorf("DB_HOST not set")
	}
	envVars.DbHost = host
	dbname, exists := os.LookupEnv("DB_NAME")
	if !exists {
		return envVars, fmt.Errorf("DB_NAME not set")
	}
	envVars.DbName = dbname
	clerkSecretKey, exists := os.LookupEnv("CLERK_SECRET_KEY")
	if !exists {
		return envVars, fmt.Errorf("CLERK_SECRET_KEY not set")
	}
	envVars.ClerkSecretKey = clerkSecretKey
	return envVars, nil
}
