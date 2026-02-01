package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port string
	Env  string // "development", "production", "test"

	// Database
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	MaxTitleLength       int
	MaxMessageLength     int
	DefaultMessagesLimit int
	MaxMessagesLimit     int
}

func Load() *Config {
	return &Config{
		// Server
		Port: getEnv("PORT", "8512"),
		Env:  getEnv("ENV", "development"),

		// Database
		DBHost:     getEnv("DB_HOST", "db"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "chatdb"),
		DBSSLMode:  getEnv("DB_SSL_MODE", "disable"),

		// Application settings
		MaxTitleLength:       getEnvAsInt("MAX_TITLE_LENGTH", 200),
		MaxMessageLength:     getEnvAsInt("MAX_MESSAGE_LENGTH", 5000),
		DefaultMessagesLimit: getEnvAsInt("DEFAULT_MESSAGES_LIMIT", 20),
		MaxMessagesLimit:     getEnvAsInt("MAX_MESSAGES_LIMIT", 100),
	}
}

// PostgresDSN - формирование dsn строки
func (c *Config) PostgresDSN() string {
	return "host=" + c.DBHost +
		" port=" + c.DBPort +
		" user=" + c.DBUser +
		" password=" + c.DBPassword +
		" dbname=" + c.DBName +
		" sslmode=" + c.DBSSLMode
}

// getEnv получает переменную окружения или значение по умолчанию
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt получает переменную окружения как int
func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
