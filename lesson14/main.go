package main

import (
	"context"
	"fmt"
	"time"
)

// context 上下文

var name string

func main () {
	// ctx, cancel := context.WithCancel(context.Background())
	// d := time.Now().Add(time.Second * 2)
	// ctx, cancel := context.WithDeadline(context.Background(), d)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 2)
	valueCtx := context.WithValue(ctx, name, "lucy")
	go work(valueCtx)
	time.Sleep(time.Second * 3)
	fmt.Println("停止工作。")
	// 主动发送取消信号
	cancel()
	time.Sleep(time.Second * 5)
}

func work(ctx context.Context) {
	for {
		select {
		case <- ctx.Done():
			fmt.Println(ctx.Value(name), "工作结束！")
			return
		default:
			fmt.Println(ctx.Value(name), "工作中。")
			time.Sleep(time.Second * 1)
		}
	}
}

