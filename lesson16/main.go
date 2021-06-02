package main

import (
	"fmt"
)

// 值类型和指针类型

type User struct {
	name string
}

func (u User) SetNameValueType(str string) {
	fmt.Printf("SetNameValueType() pointer:%p\n", &u) // SetNameValueType() pointer:0xc000096240
	u.name = str
}

func (u *User) SetNamePointerType(str string) {
	fmt.Printf("SetNamePointerType() pointer:%p\n", u) // SetNamePointerType() pointer:0xc000096220
	u.name = str
}

func (u User) ValueSetName(str string) User {
	u.name = str
	return u
}

func main() {
	user1 := &User{}
	fmt.Printf("pointer:%p\n", user1) // pointer:0xc000096220
	fmt.Println(user1)                // &{}
	user1.SetNameValueType("lucy")
	fmt.Println(user1) // &{}
	user1.SetNamePointerType("lily")
	fmt.Println(user1) // &{lily}

	// m := make(map[int]int)
	m := map[int]int{}
	fmt.Printf("map pointer:%p\n", m) // map pointer:0xc000100180
	m[0] = 1
	fmt.Printf("map pointer:%p\n", m) // map pointer:0xc000100180
	m[1] = 2

	s := make([]int, 0, 1)
	fmt.Printf("slice pointer:%p\n", s) // slice pointer:0xc00001c0a0
	s = append(s, 1)
	fmt.Printf("slice pointer:%p\n", s) // slice pointer:0xc00001c0a0
	s = append(s, 2)
	fmt.Printf("slice pointer:%p\n", s) // slice pointer:0xc00001c0b0

	user2 := &User{}
	fmt.Printf("user2 pointer:%p\n", user2) // user2 pointer:0xc000010290
	user2.SetNameValueType("tom")           // SetNameValueType() pointer:0xc0000102a0

	user3 := &User{}
	fmt.Printf("user3 pointer:%p\n", user3) // user3 pointer:0xc0000102b0
	user3.ValueSetName("bob")
	fmt.Printf("pointer:%p\n", user3) // pointer:0xc0000102b0
}
