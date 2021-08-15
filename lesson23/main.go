package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	// 连接 redis-server
	// 创建连接
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Printf("redis.Dial() error:%v", err)
		return
	}
	// 关闭连接
	defer c.Close()

	// 简单字符串
	// string
	stringExists(c)
	// stringSet(c)
	// stringSetNx(c)
	// stringSexEx(c)
	// stringStrLen(c)
	// stringSetRange(c)
	// stringAppend(c)
	// stringGet(c)

	// 结构体
	// structAdd(c)
	// structValues(c)

	// map
	// mapAdd(c)
	// mapValue(c)

	// slice
	// sliceAdd(c)
	// sliceValue(c)
}

// string
func stringSet(conn redis.Conn) {
	replySet, err := conn.Do("SET", "key1", "value1")
	if err != nil {
		fmt.Println("SET error: ", err)
	}
	fmt.Println(replySet)
}

func stringExists(conn redis.Conn) {
	// 助手函数 redis.Bool()
	replyExists, err := redis.Bool(conn.Do("EXISTS", "key1"))
	if err != nil {
		fmt.Println("EXITS error: ", err)
	}

	fmt.Println(replyExists)
}

func stringSetNx(conn redis.Conn) {
	replySetNx, err := conn.Do("SETNX", "key2", "value2")
	if err != nil {
		fmt.Println("SETNX error: ", err)
	}
	fmt.Println(replySetNx)
}

func stringSexEx(conn redis.Conn) {
	replySetNx, err := conn.Do("SETEX", "key3", 10, "value3")
	if err != nil {
		fmt.Println("SETNX error: ", err)
	}
	fmt.Println(replySetNx)
}

func stringStrLen(conn redis.Conn) {
	replyStrLen, err := conn.Do("STRLEN", "key1")
	if err != nil {
		fmt.Println("STRLEN error: ", err)
	}
	fmt.Println(replyStrLen)
}

func stringSetRange(conn redis.Conn) {
	replySetRange, err := conn.Do("SETRANGE", "key1", 5, "s")
	if err != nil {
		fmt.Println("SETRANGE error: ", err)
	}
	fmt.Println(replySetRange)
}

func stringAppend(conn redis.Conn) {
	replyAppend, err := conn.Do("APPEND", "key1", "golang")
	if err != nil {
		fmt.Println("APPEND error: ", err)
	}
	fmt.Println(replyAppend)
}

func stringGet(conn redis.Conn) {
	// 助手函数 redis.String()
	replyGet, err := redis.String(conn.Do("GET", "key1"))
	if err != nil {
		fmt.Println("GET error: ", err)
	}
	fmt.Println(replyGet)
}

type User struct {
	ID   int64  `redis:"id"`
	Name string `redis:"name"`
}

// struct
func structAdd(conn redis.Conn) {
	u1 := User{
		ID:   1,
		Name: "name1",
	}

	replyStruct, err := conn.Do("HMSET", redis.Args{}.Add("hkey1").AddFlat(&u1)...)
	if err != nil {
		fmt.Println("struct err: ", err)
	}
	fmt.Println(replyStruct)
}

func structValues(conn redis.Conn) {
	v, err := redis.Values(conn.Do("HGETALL", "hkey1"))
	if err != nil {
		fmt.Println("redis.Values() err: ", err)
	}

	// redis.ScanStruct()
	u2 := new(User)
	if err := redis.ScanStruct(v, u2); err != nil {
		fmt.Println("redis.ScanStruct() err: ", err)
	}

	fmt.Printf("%+v\n", u2)
}

// map
func mapAdd(conn redis.Conn) {
	// 注意 key 大小写
	m1 := map[string]interface{}{
		"id":   2,
		"name": "name2",
	}
	replyMap, err := conn.Do("HMSET", redis.Args{}.Add("hkey2").AddFlat(m1)...)
	if err != nil {
		fmt.Println("map err: ", err)
	}
	fmt.Println(replyMap)
}

func mapValue(conn redis.Conn) {
	v, err := redis.Values(conn.Do("HGETALL", "hkey2"))
	if err != nil {
		fmt.Println("redis.Values() err: ", err)
	}
	// redis.ScanStruct()
	u3 := new(User)
	if err := redis.ScanStruct(v, u3); err != nil {
		fmt.Println("redis.ScanStruct() err: ", err)
	}
	fmt.Printf("%+v\n", u3)
}

// slice
func sliceAdd(conn redis.Conn) {
	s := []User{
		{
			3,
			"name3",
		},
	}

	replySlice, err := conn.Do("HMSET", redis.Args{}.Add("hkey3").AddFlat(s[0])...)
	if err != nil {
		fmt.Println("slice err: ", err)
	}
	fmt.Println(replySlice)
}

func sliceValue(conn redis.Conn) {
	v, err := redis.Values(conn.Do("HGETALL", "hkey3"))
	if err != nil {
		fmt.Println("redis.Values() err: ", err)
	}
	// redis.ScanStruct()
	u4 := new(User)
	if err := redis.ScanStruct(v, u4); err != nil {
		fmt.Println("redis.ScanStruct() err: ", err)
	}
	fmt.Printf("%+v\n", u4)
}
