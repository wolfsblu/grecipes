package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"slices"
	"strings"
)

var requiredVariables = []string{
	"COOKIE_HASH_KEY",
	"COOKIE_BLOCK_KEY",
	"DB_PATH",
	"HOST",
}

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

	ensureRequiredVariables()
}

func MustGet(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("env variable '%s' is required\n", key)
	}
	return value
}

func ensureRequiredVariables() {
	var missing []string
	for _, variable := range requiredVariables {
		if _, ok := os.LookupEnv(variable); !ok {
			missing = append(missing, variable)
		}
	}
	if len(missing) > 0 {
		slices.Sort(missing)
		log.Fatalln("missing required env variable(s):", strings.Join(missing, ", "))
	}
}
