package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// 字符串性能优化

func main() {
	// 数据结构

	// string 传参

	// 字符串是只读的，不可修改的
	str := "golang"
	fmt.Println(str) // golang
	byteSlice := []byte(str)
	byteSlice[0] = 'a'
	fmt.Println(string(byteSlice)) // alang
	fmt.Println(str)               // golang

	// var str2 string = "golang"
	// fmt.Println(str2) // golang
	// ptr := (*uintptr)(unsafe.Pointer(&str2))
	// var arr *[6]byte = (*[6]byte)(unsafe.Pointer(*ptr))
	// var len *int = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&str2)) + unsafe.Sizeof((*uintptr)(nil))))
	// for i := 0; i < (*len); i++ {
	// 	fmt.Printf("%p => %c\n", &((*arr)[i]), (*arr)[i])
	// 	ptr2 := &((*arr)[i])
	// 	val := (*ptr2)
	// 	(*ptr2) = val + 1
	// }
	// fmt.Println(str)

	// 字符串和字节切片转换
	str3 := "golang"
	fmt.Printf("str3 val:%s type:%T\n", str3, str3)
	str3Ptr := (*reflect.SliceHeader)(unsafe.Pointer(&str3))
	str3Ptr.Cap = str3Ptr.Len
	fmt.Println(str3Ptr.Data)
	str4 := *(*[]byte)(unsafe.Pointer(str3Ptr))
	// str4[0] = 'a'
	fmt.Printf("str4 val:%s type:%T\n", str4, str4)
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&str4)).Data)

	// 字符串拼接
	// var b strings.Builder
}
