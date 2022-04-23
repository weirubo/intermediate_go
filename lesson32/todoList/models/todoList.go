package models

import "context"

type Todolist struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Status  int    `json:"status"`
	Created int    `json:"created"`
	Updated int    `json:"updated"`
}

type TodoListRepository interface {
	Create(ctx context.Context, t *Todolist) (err error)
}

type TodoListUsecase interface {
	Create(context.Context, *Todolist) (err error)
}
