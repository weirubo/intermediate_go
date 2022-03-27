package user

import (
	"context"
	"errors"
)

// IUser 定义接口
type IUser interface {
	Register(ctx context.Context, username, email, password string) error
}

// User 实现接口
type User struct{}

func (u User) Register(ctx context.Context, username, email, password string) error {
	if username != "" && email != "" && password != "" {
		return nil
	}
	return errors.New("register param is invalid")
}
