package config

import (
	"github.com/google/uuid"
	"os"
)

func FetchEnvValue(key string, defaultValue string) string {
	dataSourceName := os.Getenv(key)
	if dataSourceName == "" {
		dataSourceName = defaultValue
	}
	return dataSourceName
}

func GenerateUUID() (string, error) {
	uuidObj, err := uuid.NewUUID()
	data := []byte("wnw8olzvmjp0x6j7ur8vafs4jltjabi0")
	uuidObj2 := uuid.NewMD5(uuidObj, data)
	return uuidObj2.String(), err
}
