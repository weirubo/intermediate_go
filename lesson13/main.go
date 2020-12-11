package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// 修改类型的原子操作
	// 加载存储类型的原子操作
	// 相同代码在不同 CPU 架构中编译的结果不同

	// 方法：AddXXX、CompareAndSwapXXX、SwapXXX、LoadXXX、StoreXXX
	// Add 相加
	// CompareAndSwap 比较原值替换为新值
	// Swap 原值替换为新值，并返回原值
	// Load 取值
	// Store 存值

	// Value 类型 Load 和 Store

	// 多个 goroutine 的原子计数器
	var counter uint64
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100000; j++ {
				// 相加
				atomic.AddUint64(&counter, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("counter:", counter)
}
