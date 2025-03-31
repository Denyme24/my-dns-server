package env

import (
	"log"
	"os"
	"strconv"
)

func GetStringEnv(key, fallback string) string {

	val, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return val
}

func GetIntEnv(key string, fallback int) int {

	val, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}

	valInt, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("Error parsing ENV key %s : %s", key, err.Error())

	}
	return valInt
}
