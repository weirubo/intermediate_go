package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type User struct {
	Id   int64
	Name string
}

func (u *User) Register(args *User, reply *string) error {
	*reply = fmt.Sprintf("Hello %s!", args.Name)
	return nil
}

func main() {
	user := new(User)
	rpc.Register(user)
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}
}
