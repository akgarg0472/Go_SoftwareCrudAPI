package utils

import (
	"github.com/google/uuid"
	"os"
)

// GetEnvVariable utility function to retrieve environment variable or else return default value
func GetEnvVariable(key string, defaultValue string) string {
	value, found := os.LookupEnv(key)

	if !found {
		return defaultValue
	}

	return value
}

// GenerateRandomId function to generate a random string id
func GenerateRandomId() string {
	newUUID, err := uuid.NewUUID()

	if err != nil {
		return ""
	}

	return newUUID.String()
}

func IsStringValid(val *string) bool {
	return val != nil && *val != ""
}
