package main

import (
	"fmt"
	"sync"
	"time"
)

// RWMutex

type Counter struct {
	rw sync.RWMutex
	count uint64
}

func (c *Counter) Incr() {
	c.rw.Lock()
	c.count++
	c.rw.Unlock()
}

func (c *Counter) Count() uint64 {
	c.rw.RLock()
	defer c.rw.RUnlock()
	return c.count
}
func main() {
	var counter Counter

	// writer
	for i := 0; i < 10000; i++ {
		go func() {
			counter.Incr()
			time.Sleep(time.Second)
		}()
	}

	// reader
	for i := 0; i < 100000; i++ {
		go func() {
			fmt.Println(counter.count)
			time.Sleep(time.Millisecond)
		}()
	}
	fmt.Println("done!")
}
