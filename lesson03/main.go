package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
	// Open 函数需要两个参数，一个是数据库名称，另一个是 dataSourceName
	// 不建立数据库连接，仅验证参数是否有效
	// 支持数据库连接池
	if db == nil {
		db, _ = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/blog")
	}

	// 验证 dataSourceName 是否合法，验证数据库连接是否有效，必要时会创建数据库连接
	err := db.Ping()
	if err != nil {
		fmt.Printf("db.Ping() err:%v\n", err)
		return
	}
	// 限制最大开启连接数，如果 n<=0，代表不限制开启连接数，默认值为 0，不限制；如果 n 小于最大空闲连接数，将会把最大空闲连接数减小到与最大开启连接数一致。
	db.SetMaxOpenConns(10)
	// 限制最大空闲连接数，如果 n<=0，代表不保留空闲连接，当前版本默认值为 2；如果 n 大于最大开启连接数（最大开启连接数大于 0），将会把新的最大空闲连接数减小到与最大开启连接数一致。
	db.SetMaxIdleConns(5)
}

// 增
func Add() {
	// 1.写 sql
	sqlStr := "INSERT INTO user(username,email) VALUES (?, ?)"
	// 2.预处理
	// Prepare为以后的查询或执行创建一个预声明。可以从返回的预声明中并发运行多个查询或执行。当不再需要该预声明时，调用方必须调用该预声明的Close方法。
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("db.Prepare() err:%v\n", err)
		return
	}
	// 关闭预声明
	defer stmt.Close()
	user := []struct {
		username string
		email    string
	}{
		{"gopher", "gopher@99.com"},
		{"phper", "phper@99.com"},
	}
	// Exec用给定的参数执行一个预声明，并返回一个结果集。
	for _, val := range user {
		result, err := stmt.Exec(val.username, val.email)
		if err != nil {
			fmt.Printf("stmt.Exec() err:%v\n", err)
			return
		}
		// LastInsertId返回数据库的自增主键。
		id, _ := result.LastInsertId()
		fmt.Printf("ID:%d\n", id)
	}
}

// 删
func Delete() {
	// 1.写 sql
	sqlStr := "DELETE FROM user WHERE id = ?"
	// 2.预处理
	stmt, _ := db.Prepare(sqlStr)
	// 关闭预声明
	defer stmt.Close()
	// 3.执行
	result, _ := stmt.Exec(1)
	// 4.影响行数
	row, _ := result.RowsAffected()
	fmt.Printf("row:%d\n", row)
}

// 改
func Update() {
	// 1.写 sql
	sqlStr := "UPDATE user SET username=?, email=? WHERE id=?"
	// 2.预处理
	stmt, _ := db.Prepare(sqlStr)
	// 关闭预声明
	defer stmt.Close()
	// 3.执行
	result, _ := stmt.Exec("java", "java@88.com", 4)
	// 4.影响行数
	row, _ := result.RowsAffected()
	fmt.Printf("row:%d\n", row)
}

// 查单条
func SelectOne() {
	// 1.写 sql
	sqlStr := "SELECT username, email FROM user WHERE id = ?"
	// 2.预处理
	stmt, _ := db.Prepare(sqlStr)
	// 关闭预声明
	defer stmt.Close()
	// 3.查询单条
	row := stmt.QueryRow(20)
	// 4.将查询结果保存到变量
	var username, email string
	row.Scan(&username, &email)
	fmt.Printf("username:%s email:%s\n", username, email)
}

// 查多条
func Select() {
	// 1.写 sql
	sqlStr := "SELECT username,email FROM user WHERE id > ? LIMIT ?,?"
	// 2.预处理
	stmt, _ := db.Prepare(sqlStr)
	// 关闭预声明
	defer stmt.Close()
	// 3.查询多条
	rows, _ := stmt.Query(15, 0, 3)
	// 4.将查询结果保存到变量
	var username, email string
	for rows.Next() {
		rows.Scan(&username, &email)
		fmt.Printf("username:%s email:%s\n", username, email)
	}
}

// 事务
func Tx() {
	tx, _ := db.Begin()
	// 1.写 sql
	sqlStr := "UPDATE user SET username=? WHERE id=?"
	// 2.预处理
	stmt, _ := tx.Prepare(sqlStr)
	// 关闭预声明
	defer stmt.Close()
	// 3.执行
	result, _ := stmt.Exec("apple1", 4)
	// 4.影响行数
	row, _ := result.RowsAffected()

	result2, _ := stmt.Exec("java1", 6)
	row2, _ := result2.RowsAffected()

	if row > 0 && row2 > 0 {
		tx.Commit()
		fmt.Println("修改成功")
	} else {
		tx.Rollback()
		fmt.Println("修改失败")
	}

}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic() err:%v\n", err)
			return
		}
	}()
	InitDB()
	// Add()
	// Delete()
	// Update()
	// SelectOne()
	// Select()
	Tx()
}

/**
CREATE TABLE `user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(60) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `email` varchar(60) COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci
*/
