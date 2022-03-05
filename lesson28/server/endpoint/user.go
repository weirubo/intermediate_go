package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/weirubo/intermediate_go/lesson28/server/service"
)

type Request struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Name string `json:"name"`
}

func LoginEndpoint(user service.IUser) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		if err != nil {
			return nil, err
		}
		r := request.(Request)
		data := user.Login(r.Email, r.Password)
		return Response{Name: data}, nil
	}
}
