package main

import (
	"context"
	"fmt"
	"time"
)

// 并发编程 - context

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// 停止一个 goroutine
	// go func(ctx context.Context) {
	// 	for {
	// 		select {
	// 			case <-ctx.Done():
	// 				fmt.Println("goroutine 已停止")
	// 				return
	// 		default:
	// 			fmt.Println("goroutine 正在运行")
	// 			time.Sleep(time.Second)
	// 		}
	// 	}
	// }(ctx)

	// 停止多个 goroutine
	// go worker(ctx, "节点一")
	// go worker(ctx, "节点二")
	// go worker(ctx, "节点三")

	// 传递上下文信息
	ctxValue := context.WithValue(ctx, "uid", 1)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println(ctx.Value("uid"), "goroutine 已停止")
				return
			default:
				fmt.Println("goroutine 正在运行")
				time.Sleep(time.Second)
			}
		}
	}(ctxValue)
	time.Sleep(time.Second * 5)
	cancel()
	time.Sleep(time.Second * 5)
	fmt.Println("main goroutine 已结束")
}

func worker(ctx context.Context, node string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(node, "goroutine 已停止")
			return
		default:
			fmt.Println(node, "goroutine 正在运行")
			time.Sleep(time.Second)
		}
	}
}
