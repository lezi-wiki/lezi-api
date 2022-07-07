package main

import (
	"github.com/lezi-wiki/lezi-api/routers"
)

func main() {
	r := routers.InitRouter()

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
