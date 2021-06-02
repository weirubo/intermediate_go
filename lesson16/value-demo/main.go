package main

import (
	"fmt"
	"time"
)

type User struct {
	name     string
	nickName string
}

func (u *User) SetName(str string) {
	u.name = str
}

func (u *User) SetNickName(str string) {
	u.nickName = str
}

func (u *User) PrintUserInfo() {
	fmt.Printf("name:%s nickName:%s\n", u.name, u.nickName)
}

// func (u User) SetName (str string) User {
// 	u.name = str
// 	return u
// }
//
// func (u User) SetNickName (str string) User {
// 	u.nickName = str
// 	return u
// }
//
// func (u User) PrintUserInfo () {
// 	fmt.Printf("name:%s nickName:%s\n", u.name, u.nickName)
// }

func main() {
	// user := &User{}
	// fmt.Printf("user pointer:%p\n", user)
	// for i := 0; i < 10; i++ {
	// 	user.SetName(fmt.Sprintf("name-%d", i))
	// 	user.SetNickName(fmt.Sprintf("nickName-%d", i))
	// 	user.PrintUserInfo()
	// }
	for i := 0; i < 10; i++ {
		user := &User{}
		fmt.Printf("user-%d pointer:%p\n", i, user)
		go func(i int) {
			user.SetName(fmt.Sprintf("name-%d", i))
			user.SetNickName(fmt.Sprintf("nickName-%d", i+1))
			user.PrintUserInfo()
			// user.SetName(fmt.Sprintf("name-%d", i)).SetNickName(fmt.Sprintf("nickName-%d", i + 1)).PrintUserInfo()
		}(i)
	}
	time.Sleep(1 * time.Second)
}
