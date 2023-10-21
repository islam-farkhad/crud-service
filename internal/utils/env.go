package utils

import "os"

// GetAPIPort retrieves the API port from the environment variable "api_port".
// If the variable is not set, it defaults to ":9000".
func GetAPIPort() string {
	apiPort := os.Getenv("api_port")

	if apiPort == "" {
		apiPort = ":9000"
	}
	return apiPort
}
