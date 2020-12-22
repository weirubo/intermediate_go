package main

import (
	"fmt"
	"sync"
)

// WaitGroup

type Counter struct {
	mu    sync.Mutex
	count uint64
}

func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *Counter) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	var counter Counter
	var wg sync.WaitGroup
	wg.Add(10)                // 设置 WaitGroup 计数器的值
	for i := 0; i < 10; i++ { // 启动 10 个 goroutine，并发执行计数
		go func() {
			defer wg.Done() // 将 WaitGroup 计数器的值减 1
			counter.Incr()
		}()
	}
	wg.Wait() // 检查所有子 goroutine 是否全部结束
	fmt.Println(counter.count)
}
