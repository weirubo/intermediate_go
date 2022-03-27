package user

import (
	"context"
	"github.com/go-kit/kit/transport/grpc"
	pb "github.com/weirubo/intermediate_go/lesson30/pb/user"
)

type grpcHandler struct {
	register grpc.Handler
}

func (g *grpcHandler) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterRes, error) {
	_, res, err := g.register.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.RegisterRes), nil
}

func NewUserServer(ctx context.Context, endpoints Endpoints) pb.UserServiceServer {
	return &grpcHandler{
		register: grpc.NewServer(
			endpoints.UserEndpoint,
			DecodeRegister,
			EncodeRegister,
		),
	}
}

func DecodeRegister(ctx context.Context, req interface{}) (interface{}, error) {
	request := req.(*pb.RegisterReq)
	return RegisterReq{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}, nil
}

func EncodeRegister(_ context.Context, res interface{}) (interface{}, error) {
	response := res.(RegisterRes)
	return &pb.RegisterRes{
		Username: response.Username,
		Email:    response.Email,
	}, nil
}
