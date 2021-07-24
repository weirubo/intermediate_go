package v1

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"
	protoUser "micro-client/proto/user"
	"net/http"
)

type User struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

func (u *User) Login(ctx *gin.Context) {
	param := new(User)
	err := ctx.ShouldBind(&param)
	if err != nil {
		log.Debug(err)
	}

	client := NewUserClient()

	// rpc 调用远程服务的方法
	resp, err := client.Login(context.TODO(), &protoUser.LoginRequest{Email: param.Email, Password: param.Password})
	if err != nil {
		fmt.Println(err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": resp.Username,
	})

}

func NewUserClient() protoUser.UserService {
	// 创建服务
	service := grpc.NewService()

	// 创建客户端
	userClient := protoUser.NewUserService("go.micro.service.server", service.Client())

	return userClient
}
