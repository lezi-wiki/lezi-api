package main

import (
	"LeziAPI/Router"
	"fmt"
)

func main() {
	//if InitJSON == 1 {
	//	fmt.Println("MySQL database connect error.Please check your MySQL database and config file.")
	//} else {
	//	fmt.Printf("MySQL database connect success.Database id:%d\n", InitJSON)
	//}
	InitRouter := Router.InitRouter()
	if InitRouter == 1 {
		fmt.Println("Router Error!")
	}
}
