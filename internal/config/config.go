package config

import "os"

const (
	httpServerAddress = "HTTP_SERVER_ADDRESS"
)

type Config struct {
	HttpAddress string
}

func NewConfig() *Config {
	defaults := map[string]string{
		httpServerAddress: "localhost:8081",
	}

	return &Config{
		HttpAddress: grab(httpServerAddress, defaults),
	}
}

func grab(envKey string, defaults map[string]string) string {
	envVal := os.Getenv(envKey)
	if envVal != "" {
		return envVal
	}

	defaultVal, ok := defaults[envKey]
	if ok {
		return defaultVal
	}

	return ""
}
