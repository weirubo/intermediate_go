package service

import "context"

// 业务接口

type User struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

type RegisterRequest struct {
	UserName string
	Email    string
	Password string
}

type IUser interface {
	Register(ctx context.Context, req *RegisterRequest) (*User, error)
	Login(ctx context.Context, email, password string) (*User, error)
}

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) Register(ctx context.Context, req *RegisterRequest) (*User, error) {
	// TODO::操作数据库
	user := &User{
		UserName: req.UserName,
		Email:    req.Email,
	}
	return user, nil
}

func (u *UserService) Login(ctx context.Context, email, password string) (*User, error) {
	// TODO::操作数据库
	return nil, nil
}
