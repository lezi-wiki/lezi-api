package conf

import (
	"os"
	"strconv"
)

func env(name string, defaultValue string) string {
	env := os.Getenv(name)
	if env == "" {
		return defaultValue
	}
	return env
}

func envNum(name string, defaultValue int) int {
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

var SystemConfig = &system{
	Listen: env("LISTEN", ":8080"),
	Debug:  env("DEBUG", "false") == "true",
}

var DataSourceConfig = &datasource{
	Driver:   env("DB_DRIVER", "sqlite3"),
	Host:     env("DB_HOST", "localhost"),
	Port:     envNum("DB_PORT", 3306),
	Database: env("DB_DATABASE", "leziapi"),
	Username: env("DB_USERNAME", "root"),
	Password: env("DB_PASSWORD", "root"),
	File:     env("DB_FILE", "leziapi.db"),
	Prefix:   env("DB_PREFIX", "lezi_"),
	SSL:      env("DB_SSL", "false") == "true",
}

var RedisConfig = &redis{
	Network:  env("REDIS_NETWORK", "tcp"),
	Server:   env("REDIS_SERVER", ""),
	Password: env("REDIS_PASSWORD", ""),
	DB:       envNum("REDIS_DB", 0),
}
