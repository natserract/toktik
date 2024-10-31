package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/natserract/toktik/pkg/env"
)

type Configuration struct {
	Port string
	Host string

	// Third party
	RapidApiKey string
	RapiApiHost string
}

func GetConfig() *Configuration {
	err := godotenv.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load .env %v\n", err)
		os.Exit(1)
	}

	port, err := env.GetEnv("PORT")
	if err != nil {
		return nil
	}

	rapidApiKey, err := env.GetEnv("RAPID_API_KEY")
	if err != nil {
		return nil
	}

	rapidApiHost, err := env.GetEnv("RAPID_API_HOST")
	if err != nil {
		return nil
	}

	return &Configuration{
		Port:        port,
		Host:        "localhost",
		RapidApiKey: rapidApiKey,
		RapiApiHost: rapidApiHost,
	}
}
