package config

import "os"

const (
	httpServerAddress = "HTTP_SERVER_ADDRESS"
)

type Config struct {
	HttpAddress string
}

// MustLoad должна либо инстанцировать Config, либо упасть в панику
// Приставка Must значит либо функция выполнится, либо упадет в панику
// Паника здесь не реализована, но, например, в случае чтения из файла, если файл не найден - будем паниковать
func MustLoad() *Config {
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
