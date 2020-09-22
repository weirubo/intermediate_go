package main

import (
	"fmt"
	"time"
)

func main() {
	// time 包
	// 声明 Time 类型
	var t time.Time
	fmt.Println(t)
	// time.Now()函数，获取当前时间
	t1 := time.Now()
	fmt.Printf("%T\n", t1)
	fmt.Println(t1)

	/*
	 * 通过纳秒时间戳创建时间变量
	 * param: 秒，纳秒
	 * return:time.Time
	 */
	t2 := time.Unix(0, t1.UnixNano())
	fmt.Println(t2)

	// 根据自己要求创建时间
	t3 := time.Date(2020, 6, 10, 16, 51, 03, 123, time.Local)
	fmt.Println(t3)
	// 其它函数
	year := t3.Year()
	month := t3.Month()
	monthInt := int(month)
	day := t3.Day()
	hour := t3.Hour()
	minute := t3.Minute()
	second := t3.Second()
	nanosec := t3.Nanosecond()
	fmt.Println(year)
	fmt.Println(month)
	fmt.Println(monthInt)
	fmt.Println(day)
	fmt.Println(hour)
	fmt.Println(minute)
	fmt.Println(second)
	fmt.Println(nanosec)
	y, m, d := t3.Date()
	fmt.Println(y, m, d)
	fmt.Println(t3.Clock())
	// 时间类型转换字符串
	timeStr := t1.Format("2006-01-02 15:04:05")
	fmt.Printf("%T\n", timeStr)
	fmt.Printf("%s\n", timeStr)
	// 字符串类型转换成时间类型
	str := "1990-11-16 07:30:08"
	t4, _ := time.Parse("2006-01-02 15:04:05", str)
	fmt.Printf("%T\n", t4)
	fmt.Println(t4)
}
