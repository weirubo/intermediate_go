package user

import (
	"context"
	protoUser "micro-service/proto/user"
)

type User struct{}

func (u *User) Login(ctx context.Context, req *protoUser.LoginRequest, rsp *protoUser.LoginResponse) error {
	if req.Email != "gopher@88.com" || req.Password != "123456" {
		rsp.Username = "Sorry " + req.Email
		return nil
	}
	rsp.Username = "Welcome " + req.Email
	return nil
}
