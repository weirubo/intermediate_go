package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"log"
	"micro-service/handler/user"
	protoUser "micro-service/proto/user"
	"time"
)

func main() {
	// 创建服务
	service := grpc.NewService(
		micro.Name("go.micro.srv.demo"),
		micro.Version("v0.0.0"),
		micro.RegisterTTL(time.Second * 10),
		micro.RegisterInterval(time.Second * 5),
	)
	// 注册处理器
	err := protoUser.RegisterUserHandler(service.Server(), new(user.User))
	if err !=nil {
		log.Fatal(err)
	}
	// 运行服务
	if err = service.Run(); err != nil {
		log.Fatal(err)
	}
}
