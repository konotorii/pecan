package util

import (
	"os"
	"strconv"
)

type Config struct {
	Account  string
	Repo     string
	GitToken string
	Url      string
	Owner    string
	Interval int64
}

func ConfigNew() *Config {
	return &Config{
		Account:  getEnv("ACCOUNT", ""),
		Repo:     getEnv("REPOSITORY", ""),
		GitToken: getEnv("GIT_TOKEN", ""),
		Url:      getEnv("URL", ""),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}
