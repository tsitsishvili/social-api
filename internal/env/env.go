package env

import (
	"os"
	"strconv"
	"time"
)

func GetString(key, fallback string) string {
	val := os.Getenv(key)

	if val == "" {
		return fallback
	}

	return val
}

func GetInt(key string, fallback int) int {
	val := os.Getenv(key)

	if val == "" {
		return fallback
	}

	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}

	return valAsInt
}

func GetDuration(key string, fallback time.Duration) time.Duration {
	val := os.Getenv(key)

	if val == "" {
		return fallback
	}

	valAsDuration, err := time.ParseDuration(val)
	if err != nil {
		return fallback
	}

	return valAsDuration
}
