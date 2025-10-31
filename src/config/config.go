package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DBURL    = ""
	CacheUrl = ""
)

func LoadEnv() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal("Error when load .env file")
	}

	DBURL = fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL"),
	)

	CacheUrl = os.Getenv("CACHE_URL")
}
