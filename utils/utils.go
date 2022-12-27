package utils

import (
	"net/http"
	"os"
)

func UpdateHeader(toUpdatedHeader http.Header, otherHeader http.Header) {
	for key, values := range otherHeader {
		for _, value := range values {
			toUpdatedHeader.Add(key, value)
		}
	}
}

func GetVersion(defaultValue string) string {
	envValue := os.Getenv("VERSION")

	if len(envValue) == 0 {
		return defaultValue
	}

	return envValue
}
