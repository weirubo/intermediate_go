package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/weirubo/intermediate_go/lesson29/service"
)

// 接收请求
// 构建 RegisterEndpoint 和 LoginEndpoint
// 请求参数转化为 IUser 接口可以处理的参数
// 返回结果给 transport 包

type Endpoints struct {
	RegisterEndpoint endpoint.Endpoint
	LoginEndpoint    endpoint.Endpoint
}

type RegisterRequest struct {
	UserName string
	Email    string
	Password string
}

type RegisterResponse struct {
	User *service.User
}

// MakeRegisterEndpoint 构建 RegisterEndpoint
// endpoint 负责接收请求，处理请求，返回结果
// endpoint.Endpoint 是函数类型，封装一层，可以类似使用中间件，添加日志、限流、熔断、负载均衡等功能
// 也可以使用相应的 endpoint 装饰器，实现这些功能
func MakeRegisterEndpoint(iUser service.IUser) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*RegisterRequest)
		user, err := iUser.Register(ctx, &service.RegisterRequest{
			UserName: req.UserName,
			Email:    req.Email,
			Password: req.Password,
		})
		return &RegisterResponse{User: user}, err
	}
}

type LoginRequest struct {
	Email    string
	Password string
}

type LoginResponse struct {
	User *service.User
}

// MakeLoginEndpoint 构建 LoginEndpoint
func MakeLoginEndpoint(iUser service.IUser) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*LoginRequest)
		user, err := iUser.Login(ctx, req.Email, req.Password)
		return &LoginResponse{User: user}, err
	}
}
