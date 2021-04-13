package main

import (
	"fmt"
	"time"
)

type User struct {
	ID    uint
	Name  string
	Email string
}

func main() {
	u := new(User)
	fmt.Println(u)
	a := new(int)
	fmt.Println(a)
	b := new(string)
	fmt.Println(b)
	c := new(bool)
	fmt.Println(c)

	var s []int
	fmt.Println(s, len(s), cap(s))
	s = make([]int, 0, 5)
	fmt.Println(s, len(s), cap(s))

	var m map[string]int
	fmt.Println(m, len(m))
	m = make(map[string]int)
	m["a"] = 1 // 初始化后，才可以赋值
	fmt.Println(m)

	var ch chan int
	fmt.Println(len(ch), cap(ch))
	ch = make(chan int)
	go func() {
		ch <- 1
	}()
	time.Sleep(time.Second)
	fmt.Println(<-ch)
	time.Sleep(time.Second)
	fmt.Println(len(ch), cap(ch))
}
