package config

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Database struct {
	URI          string `env:"MONGO_URI" envDefault:"mongodb://localhost:27017/"`
	DatabaseName string `env:"MONGO_DATABASE" envDefault:"skeleton-app"`
}

type Origin struct {
	AllowedOrigins string `env:"ALLOWED_ORIGINS" envDefault:"*"`
}

type ServerConfig struct {
	Mode       string `env:"MODE" envDefault:"development"`
	Port       string `env:"PORT" envDefault:"8080"`
	Host       string `env:"HOST" envDefault:"localhost"`
	Database   Database
	Origin     Origin
	PrivateKey string `env:"PRIVATE_KEY"`
	PublicKey  string `env:"PUBLIC_KEY"`
}

var Config ServerConfig

func init() {
	if err := loadConfig(); err != nil {
		panic(err)
	}
}

func loadConfig() error {
	// Load environment variables into Config struct
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Error loading .env file: " + err.Error())
		return err
	}

	// Parse environment variables into Config struct
	if err := env.Parse(&Config); err != nil {
		log.Println("Error parsing environment variables: " + err.Error())
		return err
	}

	return nil
}
