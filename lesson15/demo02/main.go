package main

import "fmt"

// 组合

type User struct {
	ID    uint64
	Name  string
	Score uint64
}

type Member struct {
	User
	Level uint8
}

func (u User) Information() string {
	return fmt.Sprintf("ID:%d Name:%s Score:%d", u.ID, u.Name, u.Score)
}

func (m Member) Information() string {
	return fmt.Sprintf("ID:%d Name:%s Score:%d Level:%d", m.ID, m.Name, m.Score, m.Level)
}

func main() {
	user1 := User{
		ID:    1,
		Name:  "lucy",
		Score: 10,
	}
	user1Information := user1.Information()
	fmt.Println(user1Information)

	member1 := Member{
		User: User{
			ID:    2,
			Name:  "lily",
			Score: 51,
		},
		Level: 1,
	}
	member1Information := member1.Information()
	fmt.Println(member1Information)
}
