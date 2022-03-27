package main

import (
	"context"
	"fmt"
	pb "github.com/weirubo/intermediate_go/lesson30/pb/user"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)
	req := &pb.RegisterReq{
		Username: "gopher",
		Email:    "gopher@88.com",
		Password: "123456",
	}
	res, err := c.Register(context.Background(), req)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(res)
}
