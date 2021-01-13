package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 字符串转数值
	var str string = "100"
	/**
	 * param:string
	 * return:int,error
	 */
	num, _ := strconv.Atoi(str)
	fmt.Printf("%T\n", num)

	/*
	 * param str,进制,类型
	 * return int64,err
	 */
	num1, _ := strconv.ParseInt(str, 0, 64)
	fmt.Printf("%T\n", num1)

	// 数值转字符串
	var number int = 200
	/*
	 * param int
	 * return string
	 */
	str1 := strconv.Itoa(number)
	fmt.Printf("%T\n", str1)

	var number2 int64 = 500
	/*
	 * param int64, 进制
	 * return string
	 */
	str2 := strconv.FormatInt(number2, 10)
	fmt.Printf("%T\n", str2)
}
