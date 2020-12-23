package main

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Viper

type ServerSetting struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var server ServerSetting

func main() {
	// 实例化 viper 实例
	vp := viper.New()
	// 设置配置文件名称
	vp.SetConfigName("config")
	// 设置配置文件类型
	vp.SetConfigType("yaml")
	// 添加配置文件路径
	vp.AddConfigPath("configs/")
	// 读取配置文件内容
	err := vp.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s\n", err))
	}
	//  解码 key 到 struct
	err = vp.UnmarshalKey("Server", &server)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
	// 打印
	fmt.Printf("Server: %+v\n", server)
}
