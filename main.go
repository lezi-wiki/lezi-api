package main

import (
	_ "embed"
	"github.com/lezi-wiki/lezi-api/bootstrap"
	"github.com/lezi-wiki/lezi-api/routers"
)

//go:embed data.json
var data string

func init() {
	bootstrap.Init(data)
}

func main() {
	r := routers.InitRouter()

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
