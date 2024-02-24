package config

import "os"

const (
	HttpServerAddress = "HTTP_SERVER_ADDRESS"
	Environment       = "ENVIRONMENT"

	EnvLocal  = "local"  // Комп разработчика
	EnvServer = "server" // Сервер, дев или прод пока что не важно
)

type Config struct {
	HttpAddress string
	Environment string
}

// MustLoad должна либо инстанцировать Config, либо упасть в панику
// Приставка Must значит либо функция выполнится, либо упадет в панику
// Паника здесь не реализована, но, например, в случае чтения из файла, если файл не найден - будем паниковать
func MustLoad() *Config {
	defaults := map[string]string{
		HttpServerAddress: "localhost:8081",
		Environment:       EnvLocal,
	}

	return &Config{
		HttpAddress: grab(HttpServerAddress, defaults),
		Environment: grab(Environment, defaults),
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
