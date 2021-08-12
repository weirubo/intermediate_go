package main

import (
	"encoding/json"
	"fmt"
	"lesson24/model"
)

func main() {
	encode()
	encode2()
	decode()
	decode2()
}

// 序列化
func encode() {
	user := &model.User{
		ID:   1,
		Name: "lucy",
	}
	bs, err := user.MarshalJSON()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("encode() type=%T\nbs=%v\nstr=%s\n", bs, bs, string(bs))
}

func encode2() {
	user := model.User{
		ID:   3,
		Name: "frank",
	}
	bs, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("encode2() type=%T\nbs=%v\nstr=%s\n", bs, bs, string(bs))
}

// 反序列化
func decode() {
	user := new(model.User)
	str := `{"id":1,"name":"lucy"}`
	err := user.UnmarshalJSON([]byte(str))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("decode() user=%+v\n", user)
}

func decode2() {
	user := new(model.User)
	str := `{"id":4,"name":"bob"}`
	err := json.Unmarshal([]byte(str), user)
	if err != nil {
		fmt.Printf("err=%v\n", err)
	}
	fmt.Printf("decode2() user=%+v\n", user)
}
