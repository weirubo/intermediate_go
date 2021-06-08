package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		fmt.Println("A")
	}()
	fmt.Println("main")
	os.Exit(1)
}
