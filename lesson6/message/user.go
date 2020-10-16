package message

type User struct {
	ID   int
	Name string
}

// rpc 方法
func (u *User) GetUser(id int, user *User) error {
	userMap := map[int]User{
		1: {ID: 1, Name: "frank"},
		2: {ID: 2, Name: "lucy"},
	}
	if userInfo, ok := userMap[id]; ok {
		*user = userInfo
	}
	return nil
}
