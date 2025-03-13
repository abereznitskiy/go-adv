package env

import (
	"fmt"
	"os"
)

func Getenv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	fmt.Printf("[4-order-api]: Default value for the key %s is used: %s", key, defaultVal)

	return defaultVal
}
