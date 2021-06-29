package main

import "fmt"

func main() {
	done := make(chan struct{})
	go func() {
		fmt.Println("goroutine run over")
		done <- struct{}{}
	}()
	<-done
	fmt.Println("main goroutine run over")
}
