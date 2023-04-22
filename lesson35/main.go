package main

import "fmt"

func main() {
	//c := make(chan int) // 定义一个无缓冲区 channel
	//go func() {         // 启动一个 goroutine 调用匿名函数
	//	fmt.Println("启动一个 goroutine 调用匿名函数")
	//	c <- 1 // 该 goroutine 向 channel 发送一个值（信号）
	//}()
	//fmt.Println("main 函数")
	//out := <-c // main goroutine 从 channel 中接收一个值（信号），再未接收到值（信号）之前，一直阻塞
	//fmt.Println(out)

	//c := make(chan int, 5)
	//for i := 0; i < 20; i++ {
	//	c <- 1
	//	go func() {
	//		fmt.Println("do something:", i)
	//		<-c
	//	}()
	//}

	//c := make(chan int, 5)
	//for i := 0; i < 20; i++ {
	//	c <- 1
	//	go func(i int) {
	//		fmt.Println("do something:", i)
	//		<-c
	//	}(i)
	//}

	c := make(chan int, 5)
	for i := 0; i < 20; i++ {
		i := i
		c <- 1
		go func() {
			fmt.Println("do something:", i)
			<-c
		}()
	}

	//time.Sleep(time.Second * 2)
}
