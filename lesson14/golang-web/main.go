package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	// 多个处理器函数，可以匹配多个路由
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)
	http.HandleFunc("/client", client)
	http.HandleFunc("/param1", param1)
	http.HandleFunc("/param2", param2)
	http.HandleFunc("/json", jsonRes)

	// handler 为 nil 时，使用默认多路复用器
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "world")
}

func client(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://127.0.0.1:8080/hello")
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Fprintln(w, string(data))
}

func param1(w http.ResponseWriter, r *http.Request) {
	// http method get
	// 必须先解析，否则获取不到参数
	r.ParseForm()
	fmt.Fprintln(w, r.Form["user"][0])
}

func param2(w http.ResponseWriter, r *http.Request) {
	// http method get
	// r.FormValue() 的返回值是 string 类型
	fmt.Fprintln(w, r.FormValue("user"), r.FormValue("age"))
}

func jsonRes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ageStr := r.FormValue("age")
	ageInt, err := strconv.Atoi(ageStr)
	if err != nil {
		log.Fatal(err)
	}
	age := uint8(ageInt)
	username := r.FormValue("user")

	// 匿名（临时） struct
	jsonData, err := json.Marshal(struct {
		Username string
		Age      uint8
	}{username, age})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, string(jsonData))
}
