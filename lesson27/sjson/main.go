package main

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/sjson"
)

type User struct {
	ID    uint64
	Name  string
	Email string
}

func main() {
	u := User{
		ID:    1,
		Name:  "frank",
		Email: "frank@gmail.com",
	}
	jsonBytes, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	str := string(jsonBytes)
	fmt.Println(str)

	// sjson set a value
	set(str)
}

// sjson set a value
func set(str string) {
	s, err := sjson.Set(str, "Email", "frank@qq.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
}
