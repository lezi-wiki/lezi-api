package conf

var SystemConfig = &system{
	Listen: ":8080",
	Debug:  true,
}

var DataSourceConfig = &datasource{
	Driver:   "sqlite3",
	Host:     "localhost",
	Port:     3306,
	Database: "leziapi",
	Username: "root",
	Password: "root",
	File:     "leziapi.db",
	Prefix:   "lezi_",
}

var RedisConfig = &redis{
	Network:  "tcp",
	Server:   "",
	Password: "",
	DB:       0,
}
