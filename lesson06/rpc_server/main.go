package main

import (
	"Desktop/intermediate_go/lesson6/message"
	"net"
	"net/http"
	"net/rpc"
)

// 服务器端（调用提供方）

func main() {
	_ = rpc.Register(new(message.User))
	rpc.HandleHTTP()
	listener, _ := net.Listen("tcp", ":8081")
	_ = http.Serve(listener, nil)
}
