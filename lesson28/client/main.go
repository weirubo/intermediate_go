package main

import (
	"context"
	"fmt"
	kitEndpoint "github.com/go-kit/kit/endpoint"
	kitLog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	transportHttp "github.com/go-kit/kit/transport/http"
	consulApi "github.com/hashicorp/consul/api"
	"github.com/weirubo/intermediate_go/lesson28/client/endpoint"
	"github.com/weirubo/intermediate_go/lesson28/client/transport"
	"io"
	"log"
	"net/url"
	"os"
)

func main() {
	// 通过 Consul 连接服务端
	config := consulApi.DefaultConfig()
	config.Address = "localhost:8500"
	client, err := consulApi.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	consulClient := consul.NewClient(client)
	logger := kitLog.NewLogfmtLogger(os.Stdout)
	tags := []string{"userServer"}
	instancer := consul.NewInstancer(consulClient, logger, "userServer", tags, true)
	f := func(serverInstance string) (kitEndpoint.Endpoint, io.Closer, error) {
		tgt, _ := url.Parse("http://" + serverInstance)
		return transportHttp.NewClient("GET", tgt, transport.Req, transport.Res).Endpoint(), nil, nil
	}
	endpointer := sd.NewEndpointer(instancer, f, logger)
	endpoints, err := endpointer.Endpoints()
	if err != nil {
		log.Fatal(err)
	}
	if len(endpoints) == 0 {
		return
	}
	login := endpoints[0]
	ctx := context.Background()
	res, err := login(ctx, endpoint.Request{
		Email:    "gopher@88.com",
		Password: "123456",
	})
	if err != nil {
		log.Fatal(err)
	}
	data := res.(endpoint.Response)
	fmt.Println(data.Name)

	// 直接连服务端
	// tgt, _ := url.Parse("http://localhost:8080")
	// transportClient := transportHttp.NewClient("GET", tgt, transport.Req, transport.Res)
	// clientEndpoint := transportClient.Endpoint()
	// ctx := context.Background()
	// res, err := clientEndpoint(ctx, endpoint.Request{
	// 	Email:    "aaa",
	// 	Password: "bbb",
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// data := res.(endpoint.Response)
	// fmt.Println(data.Name)
}
