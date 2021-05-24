package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 单个处理器，不可以匹配多个路由，但是方便传参
	hello := Hello{name: "frank"}
	httpServer := http.Server{Addr: ":8080", Handler: &hello}
	httpServer.ListenAndServe()
}

type Hello struct {
	name string
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello "+h.name)
}
