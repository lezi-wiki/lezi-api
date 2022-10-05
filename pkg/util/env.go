package util

import (
	"os"
	"strconv"
)

func EnvStr(name string, defaultValue string) string {
	env := os.Getenv(name)
	if env == "" {
		return defaultValue
	}
	return env
}

func EnvNum(name string, defaultValue int) int {
	env := os.Getenv(name)
	if env == "" {
		return defaultValue
	}
	num, err := strconv.Atoi(env)
	if err != nil {
		return defaultValue
	}
	return num
}
