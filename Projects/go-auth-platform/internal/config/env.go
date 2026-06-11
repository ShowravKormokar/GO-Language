package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func LoadEnv() {

	_ = godotenv.Load("../../.env")

	accessTTL, err := time.ParseDuration(
		getEnv("JWT_ACCESS_TTL", "15m"),
	)

	if err != nil {
		log.Fatal(err)
	}

	refreshTTL, err := time.ParseDuration(
		getEnv("JWT_REFRESH_TTL", "168h"),
	)

	if err != nil {
		log.Fatal(err)
	}

	cookieSecure, _ := strconv.ParseBool(
		getEnv("COOKIE_SECURE", "false"),
	)

	AppConfig = &Config{
		AppPort: getEnv("APP_PORT", "8080"),
		AppEnv:  getEnv("APP_ENV", "development"),
		AppName: getEnv("APP_NAME", "go-auth-platform"),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBName:     getEnv("DB_NAME", "GOAuth"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),

		JWTSecret:     getEnv("JWT_SECRET", ""),
		JWTAccessTTL:  accessTTL,
		JWTRefreshTTL: refreshTTL,

		CookieSecure: cookieSecure,
		CookieDomain: getEnv("COOKIE_DOMAIN", "localhost"),

		LogLevel: getEnv("LOG_LEVEL", "debug"),
	}
}

func getEnv(key, fallback string) string {

	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}
