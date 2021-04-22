package main

import (
	"fmt"
	"time"
)

func main() {
	// block := make(chan struct{})
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println(i)
	// 	}
	// 	block <- struct{}{}
	// }()
	// go odd(block)
	// go even(block)
	// time.Sleep(time.Second)
	// <-block
	// fmt.Println("done")

	task := make(chan int, 10)
	go consumer(task)
	// 生产者
	for i := 0; i < 10; i++ {
		task <- i
	}
	time.Sleep(time.Second * 2)
}

func odd(block chan struct{}) {
	for i := 1; i <= 100; i++ {
		<-block
		if i%2 == 1 {
			fmt.Println("奇数：", i)
		}
	}
}

func even(block chan struct{}) {
	for i := 1; i <= 100; i++ {
		block <- struct{}{}
		if i%2 == 0 {
			fmt.Println("偶数：", i)
		}
	}
}

func consumer(task <-chan int) {
	for i := 0; i < 10; i++ {
		go func(id int) {
			t := <-task
			fmt.Println(id, t)
		}(i)
	}
}
