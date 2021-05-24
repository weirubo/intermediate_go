package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 创建一个多路复用器
	serveMux := http.NewServeMux()
	hello := Hello{}
	serveMux.Handle("/hello", hello)
	http.ListenAndServe(":8080", serveMux)
}

type Hello struct{}

// 创建处理器
func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}
