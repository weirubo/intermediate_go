package main

import (
	"fmt"
	"sync"
)

// Cond

func main() {
	var mu sync.Mutex
	// 创建 Cond
	cond := sync.NewCond(&mu)

	// 计数
	var count uint64

	// 报名表
	var stuSlice []int

	// 模拟学生报名参加课外活动
	for i := 0; i < 30; i++ {
		go func(i int) {
			cond.L.Lock()
			stuSlice = append(stuSlice, i)
			count++
			cond.L.Unlock()

			// 唤醒所有等待此 cond 的 goroutine
			cond.Broadcast()
		}(i)
	}

	// 调用 Wait 方法前，调用者必须持有锁
	cond.L.Lock()
	for count != 30 {
		// 调用者被阻塞，并被放入 cond 的等待队列中
		cond.Wait()
	}
	cond.L.Unlock()

	fmt.Println(len(stuSlice), stuSlice)
}
