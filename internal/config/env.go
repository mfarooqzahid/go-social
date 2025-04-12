package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	Port       string
	PublicHost string
	DBAddress  string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
}

func LoadEnv() *config {

	godotenv.Load()

	return &config{
		PublicHost: GetEnv("PUBLIC_HOST", "localhost"),
		Port:       GetEnv("PORT", "8080"),
		DBUser:     GetEnv("DB_USER", "admin"),
		DBAddress: fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			GetEnv("DB_HOST", "localhost"),
			GetEnv("DB_PORT", "5432"),
			GetEnv("DB_USER", "admin"),
			GetEnv("DB_PASSWORD", "adminpassword"),
			GetEnv("DB_NAME", "go-social"),
		),
		DBName:    GetEnv("DB_NAME", "go-social"),
		JWTSecret: GetEnv("JWT_SECRET", "my-jwt-secret"),
	}
}

var Envs = LoadEnv()

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}
