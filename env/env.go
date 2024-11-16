package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("failed to load .env file", err)
	}
}

func MustGet(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("env variable '%s' is required\n", key)
	}
	return value
}
