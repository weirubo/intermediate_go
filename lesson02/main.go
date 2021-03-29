package main

import (
	"fmt"
	"unsafe"
)

type user struct {
	name string
	age  uint
}

func main() {
	// 定义一个指针变量
	student := new(user)
	// user 结构体中的 name 字段是第一个字段，可以直接通过指针修改，不需要使用偏移
	studentName := (*string)(unsafe.Pointer(student))
	*studentName = "lucy"
	// user 结构体中的 age 字段不是第一个字段，所以需要使用偏移才能找到 age 字段的内存地址，修改值
	studentAge := (*uint)(unsafe.Pointer(uintptr(unsafe.Pointer(student)) + unsafe.Offsetof(student.age)))
	*studentAge = 18
	fmt.Println(*student)
}
