package main

import (
	"context"
	pb "github.com/weirubo/intermediate_go/lesson30/pb/user"
	"github.com/weirubo/intermediate_go/lesson30/rpc-server/user"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	ctx := context.Background()
	// var userServer user.IUser
	userServer := user.User{}
	endpoints := user.Endpoints{
		UserEndpoint: user.MakeUserEndpoint(userServer),
	}

	handler := user.NewUserServer(ctx, endpoints)
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, handler)
	err = s.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}
