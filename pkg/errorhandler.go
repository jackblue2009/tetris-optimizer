package pkg

import (
	"fmt"
	"os"
)

func FatalError(s string) {
	fmt.Println("\033[31m" + s + "\033[0m")
	os.Exit(0)
}

func ErrorMsg() {
	fmt.Println("ERROR")
	os.Exit(0)
}

func PrintUsage() {
	fmt.Println("\033[33mUSAGE: go run main.go <filename>\nEXAMPLE: go run main.go sample.txt\033[0m")
	os.Exit(0)
}
