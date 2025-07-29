package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	RedisUrl   string
	BucketName string
}

func InitConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		RedisUrl:   os.Getenv("REDIS_URL"),
		BucketName: os.Getenv("BUCKET_NAME"),
	}
}
