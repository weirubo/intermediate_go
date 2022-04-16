package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world")
		LogInfo("run success")
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}

}

func LogInfo(logs string) {
	file, err := os.OpenFile("service.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("create log file failed!")
	}
	defer file.Close()
	logger := log.New(file, "", log.LstdFlags)
	logger.Println("[info]" + logs)
}
