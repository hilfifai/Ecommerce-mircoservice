package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	APIPort       string
	DBHost        string
	DBUser        string
	DBPassword    string
	DBName        string
	DBPort        string
	DBSslMode     string
	RedisAddr     string
	RedisPassword string
	RedisDB       int
	JWTSecret     string
	N8nWebhookURL string
}

func LoadConfig(path string) (*Config, error) {
	// Coba muat file .env. Jika tidak ada, jangan panik, cukup catat sebagai peringatan.
	// Ini memungkinkan aplikasi berjalan baik secara lokal (dengan .env) maupun di Docker (tanpa .env).
	if err := godotenv.Load(path + "/.env"); err != nil {
		log.Println("Warning: .env file not found. Reading configuration from environment variables.")
	}

	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	// Ambil konfigurasi dari environment variables.
	// Di Docker, variabel ini disuntikkan oleh docker-compose.yml.
	// Secara lokal, variabel ini diisi oleh godotenv dari file .env.
	return &Config{
		APIPort:       os.Getenv("API_PORT"),
		DBHost:        os.Getenv("DB_HOST"),
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),
		DBPort:        os.Getenv("DB_PORT"),
		DBSslMode:     os.Getenv("DB_SSLMODE"), // Pastikan ini ada di .env dan docker-compose.yml
		RedisAddr:     os.Getenv("REDIS_ADDR"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RedisDB:       redisDB,
		JWTSecret:     os.Getenv("JWT_SECRET"),
		N8nWebhookURL: os.Getenv("N8N_WEBHOOK_URL"),
	}, nil // Selalu return nil error di sini untuk melanjutkan eksekusi.
}