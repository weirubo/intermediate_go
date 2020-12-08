package main

import (
	"fmt"
	"sync"
)

// Once
func main() {
	var once sync.Once
	func1 := func() {
		fmt.Println("func1")
	}
	once.Do(func1)
	func2 := func() {
		fmt.Println("func2")
	}
	once.Do(func2)
}
