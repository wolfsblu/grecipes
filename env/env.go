package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Load() {
	env, ok := os.LookupEnv("APP_ENV")
	if !ok {
		env = "dev"
	}

	_ = godotenv.Load(fmt.Sprintf(".env.%s.local", env))
	if env != "test" {
		_ = godotenv.Load(".env.local")
	}

	_ = godotenv.Load(fmt.Sprintf(".env.%s", env))
	_ = godotenv.Load()
}

func MustGet(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("env variable '%s' is required\n", key)
	}
	return value
}
