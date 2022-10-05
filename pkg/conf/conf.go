package conf

const defaultConf = `[System]
Listen = :8080
Debug = false
`

const Version = "1.1.0"

type system struct {
	Listen     string `validate:"required"`
	Debug      bool
	HashIDSalt string
}

type datasource struct {
	Driver   string `validate:"required"`
	Host     string
	Port     int
	Database string
	Username string
	Password string
	File     string
	Prefix   string
	SSLMode  string
}

type redis struct {
	Network  string
	Server   string
	Password string
	DB       int
}
