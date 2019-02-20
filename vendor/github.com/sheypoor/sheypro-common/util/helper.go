package util

import "os"

// IsDev returns true if the app is not in production mode.
func IsDev() bool {
	return GetEnv("APP_ENV", "develop") != "production"
}

// GetEnv is a convenience function for getting environment variables.
func GetEnv(varName string, defaultValue string) string {
	if value := os.Getenv(varName); value != "" {
		return value
	}
	return defaultValue
}
