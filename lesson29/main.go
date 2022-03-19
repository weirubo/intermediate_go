package main

import (
	"context"
	"github.com/weirubo/intermediate_go/lesson29/endpoint"
	"github.com/weirubo/intermediate_go/lesson29/service"
	"github.com/weirubo/intermediate_go/lesson29/transport"
	"log"
	"net/http"
)

// 应用主入口

// 如何基于 Go kit 开发一个 Web 项目

func main() {
	ctx := context.Background()
	userService := service.NewUserService()
	endpoints := &endpoint.Endpoints{
		RegisterEndpoint: endpoint.MakeRegisterEndpoint(userService),
		LoginEndpoint:    endpoint.MakeLoginEndpoint(userService),
	}

	r := transport.NewHttpHandler(ctx, endpoints)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
		return
	}
}
