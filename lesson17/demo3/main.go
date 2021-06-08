package main

import "fmt"

func main() {
	fmt.Println("main")
	return
	defer func() {
		fmt.Println("A")
	}()
}
