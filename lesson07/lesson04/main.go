package main

import (
	"fmt"
	"sync"
)

// 线程安全队列
func main() {
	// 定义队列
	queue := NewSliceQueue(5)
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			queue.Push("ping")
		}()
	}
	wg.Wait()
	fmt.Println(len(queue.data))
	for j := 0; j < 5; j++ {
		fmt.Println(queue.Pop())
	}
	fmt.Println(len(queue.data))
}

// 组合数据结构队列
type SliceQueue struct {
	data []interface{}
	mu sync.Mutex
}

// 新建队列
func NewSliceQueue(n int) (q *SliceQueue) {
	return &SliceQueue{data: make([]interface{}, 0, n)}
}

// 入队
func (s *SliceQueue) Push(val interface{}) {
	s.mu.Lock()
	s.data = append(s.data, val)
	s.mu.Unlock()
}

// 出队
func (s *SliceQueue) Pop() interface{} {
	s.mu.Lock()
	if len(s.data) == 0 {
		s.mu.Unlock()
		return nil
	}
	val := s.data[0]
	s.data = s.data[1:]
	s.mu.Unlock()
	return val
}
