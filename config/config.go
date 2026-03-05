package config

import "os"

type Config struct {
	Port     string
	DBPath   string
	AuthUser string
	AuthPass string
}

func LoadConfig() Config {

	return Config{
		Port:     getEnv("PORT", "8080"),
		DBPath:   getEnv("DB_PATH", "./todos.db"),
		AuthUser: getEnv("AUTH_USER", "admin"),
		AuthPass: getEnv("AUTH_PASS", "password"),
	}
}

func getEnv(key, defaultValue string) string {

	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}