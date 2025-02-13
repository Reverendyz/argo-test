package utils

import "os"

func GetEnvOrFallback(envName, fallBack string) string {
	if os.Getenv(envName) == "" {
		return fallBack
	}
	return os.Getenv(envName)
}
