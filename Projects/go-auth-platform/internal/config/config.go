package config

import "time"

type Config struct {
	AppPort string
	AppEnv  string
	AppName string

	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string

	JWTSecret     string
	JWTAccessTTL  time.Duration
	JWTRefreshTTL time.Duration

	CookieSecure bool
	CookieDomain string

	LogLevel string
}

var AppConfig *Config
