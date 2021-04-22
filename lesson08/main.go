package main

import (
	"fmt"
	"sync"
)

// 并发编程 - 锁

func main() {
	var mu sync.Mutex
	var a int
	mu.Lock()
	// mu.Lock() // 重复锁定
	a++
	mu.Unlock()
	// mu.Unlock() // 重复解锁
	fmt.Println(a)

	aFunc(mu)
	bFunc(mu)

	aaFunc()
	bbFunc()
}

// 值传递
func aFunc(mu sync.Mutex) {
	var a int
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			wg.Done()
			mu.Lock()
			a++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("aFunc", a)
}

func bFunc(mu sync.Mutex) {
	var a int
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			wg.Done()
			mu.Lock()
			a++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("bFunc", a)
}

func aaFunc() {
	var a int
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			wg.Done()
			mu.Lock()
			a++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("aaFunc", a)
}

func bbFunc() {
	var a int
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			wg.Done()
			mu.Lock()
			a++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("bbFunc", a)
}
