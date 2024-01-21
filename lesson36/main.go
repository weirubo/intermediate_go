package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Golang 语言开发栈")
	//go func() {
	//	fmt.Println("Golang 公众号")
	//}()
	//select {}
	//time.Sleep(time.Second)
	c := make(chan string)
	//c := make(chan string, 1)
	//c := make(chan string, 2)
	//c <- "Go"
	//DoChannel(c)
	close(c)
	DoChannelV2(c)
}

func DoChannel(c chan string) {
	var receive string
	send := "golang"
	select {
	case receive = <-c:
		fmt.Println(receive)
	case c <- send:
		fmt.Println(send)
	}
}

func DoChannelV2(c chan string) {
	var (
		receive string
		ok      bool
	)
	select {
	case receive, ok = <-c:
		if !ok {
			fmt.Println("no data")
		} else {
			fmt.Println(receive)
		}
	}
}
