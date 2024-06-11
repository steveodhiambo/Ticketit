package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                            string
	DBUser                          string
	DBPassword                      string
	DBHost                          string
	DBPort                          string
	DBName                          string
	JWTSecret                       string
	JWTExpirationInSeconds          int64
	RefreshTokenExpirationInSeconds int64
}

// Envs is a singleton that holds initial configs
var Envs = initConfig()

func initConfig() Config {
	godotenv.Load(".env")
	return Config{
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBName:     getEnv("DB_NAME", "test"),
		//DBPort:                          getEnvAsInt("DB_PORT", 5432),
		JWTSecret:                       getEnv("JWT_SECRET", "not-secret-anymore"),
		JWTExpirationInSeconds:          getEnvAsInt("JWT_EXP", 15*60),
		RefreshTokenExpirationInSeconds: getEnvAsInt("JWT_REFRESH_EXP", 3600*24*7),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}

	return fallback
}
