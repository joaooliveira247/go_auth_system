package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DBURL         = ""
	CacheUrl      = ""
	CacheDuration = 7200
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
	CacheDuration, err = strconv.Atoi(os.Getenv("CACHE_DURATION"))

	if err != nil {
		log.Fatal(err)
	}
}
