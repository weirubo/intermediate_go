package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// rpc

func main() {
	// 拨号 RPC 服务
	client, err := rpc.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal("dial error:", err)
	}

	var reply string
	// 调用 RPC 方法
	err = client.Call("PersonService.Eat", "food", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
