package main

import (
	"log"
	"micro-client/router"
	"net/http"
	"time"
)

func main() {
	r := router.NewRouter()
	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
		MaxHeaderBytes: 1 << 20,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
