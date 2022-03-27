package user

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"log"
)

type RegisterReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRes struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Endpoints struct {
	UserEndpoint endpoint.Endpoint
}

func MakeUserEndpoint(user IUser) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(RegisterReq)
		err = user.Register(ctx, req.Username, req.Email, req.Password)
		if err != nil {
			log.Printf("err:%s", err)
		}
		return RegisterRes{
			Username: req.Username,
			Email:    req.Email,
		}, nil
	}
}
