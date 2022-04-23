package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/weirubo/intermediate_go/lesson32/todoList/delivery/http"
	"github.com/weirubo/intermediate_go/lesson32/todoList/repository/mysql"
	"github.com/weirubo/intermediate_go/lesson32/todoList/usecase"
	"log"
)

func main() {
	conn, err := sql.Open(`mysql`, "root:root@tcp(127.0.0.1:3306)/todolist")
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	todoListRepository := mysql.NewMysqlTodoListRepository(conn)
	todoListUsecase := usecase.NewTodoListUsecase(todoListRepository)
	http.NewTodoListHandler(r, todoListUsecase)
}
