package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds the application configuration
type Config struct {
	APIURL   string
	APIToken string
	Port     int
}

// New creates a new Config instance with values from environment variables
func New() (*Config, error) {
	port := 3000 // default port
	if portStr := os.Getenv("PORT"); portStr != "" {
		var err error
		port, err = strconv.Atoi(portStr)
		if err != nil {
			return nil, fmt.Errorf("invalid PORT value: %v", err)
		}
	}

	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		apiURL = "http://localhost:8000" // default API URL
	}

	apiToken := os.Getenv("API_TOKEN")

	return &Config{
		APIURL:   apiURL,
		APIToken: apiToken,
		Port:     port,
	}, nil
}
