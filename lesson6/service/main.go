package main

import (
	"log"
	"net"
	"net/rpc"
)

// rpc

// 1.创建一个自定义类型
type PersonService struct {
}

// 2.创建方法
/*
RPC 规则：
1. 方法只能有两个可序列化的参数，其中第二个参数是指针类型。第一个参数代表调用者提供的参数，第二个参数代表返回给调用者的参数。
2. 返回一个 error 类型，返回值如果不是 nil，将被作为字符串回传，就和 error.New() 创建的一样。如果返回了错误，回复的参数将不会被发送给客户端。
3. 方法必须是可导出的
*/
func (p *PersonService) Eat(request string, reply *string) error {
	*reply = "person eat " + request
	return nil
}

func (p *PersonService) Sleep(request string, reply *string) error {
	*reply = "person" + request
	return nil
}

func main() {
	// 3. 注册
	/*
		1. 服务端注册一个对象（服务），对象的名字是该对象的类型名。
		2. 服务端可以注册多个不同类型的对象，但是不能注册多个相同类型的对象。
	*/
	rpc.RegisterName("PersonService", new(PersonService))

	// 4. 监听
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal("Listen tcp error:", err)
	}

	// 5. Accept()
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	// 6. ServeConn()
	rpc.ServeConn(conn)
}
