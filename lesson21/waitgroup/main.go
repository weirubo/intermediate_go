package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(id int) {
			fmt.Println(id, "运行结束")
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("main goroutine run over")
}
