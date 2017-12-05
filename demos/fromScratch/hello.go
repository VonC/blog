package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello World!")
	args := os.Args
	if len(args) > 1 {
		fmt.Printf("Actually, Hello '%s'", strings.Join(os.Args[1:], " "))
	}
}
