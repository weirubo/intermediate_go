package main

import "fmt"

func main() {
	a := 0
	defer func(num int) {
		fmt.Println("defer func()", num)
	}(a)
	a++
	fmt.Println(a)
}
