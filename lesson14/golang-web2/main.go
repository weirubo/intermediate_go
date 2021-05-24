package main

import (
	"fmt"
	"net/http"
)

func main() {
	httpServer := http.Server{
		Addr: ":8080",
	}
	// 多个处理器，可以匹配多个路由
	// 默认使用 DefaultServeMux
	http.Handle("/hello", Hello{})
	http.Handle("/world", World{})
	httpServer.ListenAndServe()
}

type Hello struct{}

type World struct{}

/*
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
*/
func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

func (w1 World) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "world")
}
