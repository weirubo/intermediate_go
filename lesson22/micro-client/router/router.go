package router

import (
	"github.com/gin-gonic/gin"
	"micro-client/router/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	user := new(v1.User)
	// 路由分组
	v1 := r.Group("/v1")
	{
		v1.POST("/login", user.Login)
	}
	return r
}
