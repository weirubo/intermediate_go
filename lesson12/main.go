package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// 拦截系统信号
// 优雅退出 http server

func main() {
	// c := make(chan os.Signal, 1)
	// 捕获给定信号，并存入 channel c
	// 如果未给定信号，将捕获所有信号，并存入 channel c
	// signal.Notify(c, os.Interrupt)
	// 阻塞
	// s := <-c
	// fmt.Println(s)

	// 优雅退出
	http.HandleFunc("/", hello)
	server := http.Server{Addr: ":8080"}
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Println("server start failed")
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	s := <-c
	fmt.Printf("接收信号：%s\n", s)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("server shutdown failed")
	}
	fmt.Println("server exit")
}

func hello(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	fmt.Fprintln(w, "Hello Go!")
}
