package conf

type system struct {
	Listen string `validate:"required"`
	Debug  bool
}

const defaultConf = `[System]
Listen = :8080
Debug = false
`
