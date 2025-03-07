package env

import (
	"os"
	"strconv"
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
