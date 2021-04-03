package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// 数组和切片有什么区别？

func main() {
	var arr1 [4]int
	fmt.Printf("arr1 val:%d arr1 len:%d arr1 cap:%d\n", arr1, len(arr1), cap(arr1))
	arr := [4]int{}
	fmt.Printf("val:%d len:%d cap:%d\n", arr, len(arr), cap(arr)) // val:[0 0 0 0] len:4 cap:4
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	// arr[4] = 5 // invalid array index 4 (out of bounds for 4-element array)
	fmt.Printf("val:%d len:%d cap:%d\n", arr, len(arr), cap(arr)) // val:[1 2 3 4] len:4 cap:4
	arr2 := arr
	fmt.Printf("arr2 val:%d len:%d cap:%d ptr:%p\n", arr2, len(arr2), cap(arr2), &arr2)
	fmt.Printf("arr val:%d len:%d cap:%d ptr:%p\n", arr, len(arr), cap(arr), &arr)
	ss := arr[:]
	ssPtr := (*reflect.SliceHeader)(unsafe.Pointer(&ss)).Data
	fmt.Printf("ss val:%d len:%d cap:%d ptr:%v\n", ss, len(ss), cap(ss), ssPtr)
	ss2 := arr[:]
	ss2Ptr := (*reflect.SliceHeader)(unsafe.Pointer(&ss2)).Data
	fmt.Printf("ss2 val:%d len:%d cap:%d ptr:%v\n", ss2, len(ss2), cap(ss2), ss2Ptr)

	var s []int
	if s == nil {
		fmt.Println("nil")
	}
	fmt.Printf("s val:%d len:%d cap:%d\n", s, len(s), cap(s))
	s = append(s, 1)
	fmt.Printf("s val:%d len:%d cap:%d\n", s, len(s), cap(s))
	s = append(s, 2)
	fmt.Printf("s val:%d len:%d cap:%d\n", s, len(s), cap(s))
	s = append(s, 3)
	fmt.Printf("s val:%d len:%d cap:%d\n", s, len(s), cap(s))
}
