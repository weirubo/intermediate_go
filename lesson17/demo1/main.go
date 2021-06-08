package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("A")
	}()

	defer func() {
		fmt.Println("B")
	}()

	fmt.Println("main goroutine run over")

	// panic("this is a panic example")

	// return
}
