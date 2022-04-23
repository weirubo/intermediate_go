package mysql

import (
	"context"
	"database/sql"
	"github.com/weirubo/intermediate_go/lesson32/todoList/models"
)

type mysqlTodoListRepository struct {
	Conn *sql.DB
}

func NewMysqlTodoListRepository(Conn *sql.DB) models.TodoListRepository {
	return &mysqlTodoListRepository{Conn}
}

func (m *mysqlTodoListRepository) Create(ctx context.Context, t *models.Todolist) (err error) {
	query := "INSERT INTO todolist(title,status) VALUES(?, ?)"
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	result, err := stmt.ExecContext(ctx, t.Title, t.Status)
	if err != nil {
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		return
	}
	t.Id = id
	return
}
