package main

import (
	"LeziAPI/routers"
)

func main() {
	r := routers.InitRouter()

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
