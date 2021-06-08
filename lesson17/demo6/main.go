package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("this is a panic")
		}
	}()
	panic("this is a test panic")
	fmt.Println("main")
}
