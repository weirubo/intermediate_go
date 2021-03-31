package main

import (
	"fmt"
)

// Golang 语言编程技巧

var a = "hello"

func main() {
	var name string
	name = "frank"
	fmt.Printf("val:%s, type:%T\n", name, name)
	// 类型推断
	name2 := "lucy"
	fmt.Printf("val:%s type:%T\n", name2, name2)

	var name3 = "lily"
	fmt.Printf("val:%s type:%T\n", name3, name3)

	// 重新声明变量
	name, age := "bob", 10
	fmt.Printf("name:%s age:%d\n", name, age)

	// 同名变量，不同作用域允许变量同名
	a := "world"
	fmt.Printf("a:%s\n", a)
}
