package config

import "os"

func FetchEnvValue(key string, defaultValue string) string {
	dataSourceName := os.Getenv(key)
	if dataSourceName == "" {
		dataSourceName = defaultValue
	}
	return dataSourceName
}
