package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "frank魏如博"
	fmt.Println(len(str))
	// 字符串截取
	str1 := str[0]
	fmt.Println(str1)
	fmt.Printf("%T\n", str1)
	str2 := fmt.Sprintf("%c", str1)
	// 取字符
	fmt.Printf("%T\n", str2)
	fmt.Println(str2)

	// 范围截取
	str3 := str[0:2]
	fmt.Printf("%T\n", str3)
	fmt.Println(str3)
	str4 := str[:3]
	fmt.Printf("%T\n", str4)
	fmt.Println(str4)
	str5 := str[5:]
	fmt.Printf("%T\n", str5)
	fmt.Println(str5)

	// 字符串转切片
	s1 := []rune(str)
	fmt.Printf("%T\n", s1)
	fmt.Println(len(s1))
	fmt.Println(s1[0])
	// 取字符
	fmt.Printf("%c\n", s1[0])
	fmt.Printf("%c\n", s1[5])

	// 遍历字符串
	for i, n := range str {
		// fmt.Println(i, n)
		// 打印字符
		fmt.Printf("%d=>%c\n", i, n)
	}

	// 常用字符串函数
	// HasPrefix 前缀
	// isOk := strings.HasPrefix(str, "fr")
	// HasSuffix 后缀
	// isOk := strings.HasSuffix(str, "博")
	// Contains 包含子串
	// isOk := strings.Contains(str, "an")
	// fmt.Printf("%t\n", isOk)
	// Index 子串第一次出现的位置
	// index := strings.Index(str, "r")
	// LastIndex 子串最后出现的位置
	index := strings.LastIndex(str, "a")
	fmt.Printf("%d\n", index)
	// ToLower 转小写
	strL := strings.ToLower(str)
	fmt.Printf("%s\n", strL)
	// ToUpper 转大写
	strU := strings.ToUpper(str)
	fmt.Printf("%s\n", strU)
	// Replace 替换
	strNew := strings.Replace(str, "a", "b", 1)
	fmt.Printf("%s\n", strNew)
	// Trim 去除字符串前后的指定字符
	strNew2 := strings.Trim(str, "f")
	fmt.Printf("%s\n", strNew2)
	// Split 字符串切分成切片
	slice := strings.Split(str, "")
	fmt.Println(slice)
	// Join 切片组合成字符串
	strSlice := strings.Join(slice, "+")
	fmt.Printf("%s\n", strSlice)

}
