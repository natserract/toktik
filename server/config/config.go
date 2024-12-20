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
	OpenAIKey   string
}

func GetConfig() *Configuration {
	// Railway doesn't need to load environment variables from an .env file, that's only for local development.
	// Do nothing
	err := godotenv.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load .env %v\n", err)
	}

	port, err := env.GetEnv("PORT")
	if err != nil {
		return nil
	}

	host, err := env.GetEnv("HOST")
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

	openAIKey, err := env.GetEnv("OPENAI_KEY")
	if err != nil {
		return nil
	}

	return &Configuration{
		Port:        port,
		Host:        host,
		RapidApiKey: rapidApiKey,
		RapiApiHost: rapidApiHost,
		OpenAIKey:   openAIKey,
	}
}
