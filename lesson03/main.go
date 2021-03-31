package main

import (
	"errors"
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

	// 类型断言
	id, err := getVal("a")
	if err != nil {
		fmt.Printf("err:%s\n", err)
		return
	}
	fmt.Println(id)

	// 类型转换

	// 类型切换
	var aa interface{}
	// aa = 1
	// aa = "golang"
	aa = false
	switch val := aa.(type) {
	case int:
		fmt.Printf("val:%d type:%T\n", val, val)
	case string:
		fmt.Printf("val:%s type:%T\n", val, val)
	default:
		fmt.Printf("unknow type:%T\n", val)
	}
}

// 类型断言
func getVal(val interface{}) (interface{}, error) {
	param, ok := val.(int)
	if !ok {
		err := errors.New("illegal parameter")
		return param, err
	}
	return param, nil
}
