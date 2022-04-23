package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/weirubo/intermediate_go/lesson32/todoList/models"
	"log"
	"net/http"
)

type TodoListHandler struct {
	TodoListUsecase models.TodoListUsecase
}

func NewTodoListHandler(r *gin.Engine, todoListUsecase models.TodoListUsecase) {
	handler := &TodoListHandler{
		TodoListUsecase: todoListUsecase,
	}
	r.POST("/create", handler.Create)
	r.Run()
}

func (t *TodoListHandler) Create(c *gin.Context) {
	var todoList models.Todolist
	err := c.Bind(&todoList)
	if err != nil {
		log.Fatal(err)
	}
	ctx := c.Request.Context()
	err = t.TodoListUsecase.Create(ctx, &todoList)
	if err != nil {
		fmt.Printf("Create() || err=%v \n", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
	return
}
