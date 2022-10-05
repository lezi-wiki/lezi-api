package conf

import (
	"github.com/lezi-wiki/lezi-api/pkg/util"
)

var SystemConfig = &system{
	Listen:     util.EnvStr("LISTEN", ":8080"),
	Debug:      util.EnvStr("DEBUG", "false") == "true",
	HashIDSalt: util.EnvStr("HASHID_SALT", ""),
}

var DataSourceConfig = &datasource{
	Driver:   util.EnvStr("DB_DRIVER", "sqlite3"),
	Host:     util.EnvStr("DB_HOST", "localhost"),
	Port:     util.EnvNum("DB_PORT", 3306),
	Database: util.EnvStr("DB_DATABASE", "leziapi"),
	Username: util.EnvStr("DB_USERNAME", "root"),
	Password: util.EnvStr("DB_PASSWORD", "root"),
	File:     util.EnvStr("DB_FILE", "leziapi.db"),
	Prefix:   util.EnvStr("DB_PREFIX", "lezi_"),
	SSLMode:  util.EnvStr("DB_SSL", "disable"),
}

var RedisConfig = &redis{
	Network:  util.EnvStr("REDIS_NETWORK", "tcp"),
	Server:   util.EnvStr("REDIS_SERVER", ""),
	Password: util.EnvStr("REDIS_PASSWORD", ""),
	DB:       util.EnvNum("REDIS_DB", 0),
}
