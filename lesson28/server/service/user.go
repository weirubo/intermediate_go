package service

type IUser interface {
	Login(email, password string) string
}

type User struct {
}

func (u User) Login(email, password string) string {
	if email != "" && password != "" {
		return "user"
	}
	return "guest"
}
