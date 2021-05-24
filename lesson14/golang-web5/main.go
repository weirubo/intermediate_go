package main

import (
	"fmt"
	"net/http"
)

func main() {
	serveMux := http.NewServeMux()
	// 将传入函数参数转义为 HandlerFunc
	hello := http.HandlerFunc(Hello)
	serveMux.Handle("/hello", hello)
	http.ListenAndServe(":8080", serveMux)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}
