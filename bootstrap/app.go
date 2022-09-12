package bootstrap

import (
	"fmt"
	"os"
)

func printName() {
	bytes, _ := os.ReadFile("bootstrap/text")
	fmt.Print(string(bytes))
}
