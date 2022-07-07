package main

import (
	"LeziAPI/Router"
	"fmt"
)

func main() {
	InitRouter := Router.InitRouter()
	if InitRouter == 1 {
		fmt.Println("Router Error!")
	}
}
