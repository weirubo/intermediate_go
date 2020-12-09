package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type User struct {
	Id   int64
	Name string
}

func main() {
	client, err := jsonrpc.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	user := User{
		Id:   1,
		Name: "lucy",
	}
	var reply string
	err = client.Call("User.Register", user, &reply)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(reply)
}
