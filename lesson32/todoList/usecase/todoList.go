package usecase

import (
	"context"
	"fmt"
	"github.com/weirubo/intermediate_go/lesson32/todoList/models"
)

type todoListUsecase struct {
	todoListRepo models.TodoListRepository
}

func NewTodoListUsecase(t models.TodoListRepository) models.TodoListRepository {
	return &todoListUsecase{
		todoListRepo: t,
	}
}

func (tl *todoListUsecase) Create(ctx context.Context, t *models.Todolist) (err error) {
	if t.Title == "" {
		return fmt.Errorf("illegal parameter")
	}
	return tl.todoListRepo.Create(ctx, t)
}
