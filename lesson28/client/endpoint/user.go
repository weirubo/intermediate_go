package endpoint

type Request struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Name string `json:"name"`
}
