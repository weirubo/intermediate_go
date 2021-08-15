package main

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
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

	// gjson get a value
	get(str)
}

// gjson get a value
func get(str string) {
	rst := gjson.Get(str, "Name")
	fmt.Println(rst)
}
