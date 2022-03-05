package main

import (
	"fmt"
	kitLog "github.com/go-kit/kit/log"
	transportHttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	consulApi "github.com/hashicorp/consul/api"
	"github.com/weirubo/intermediate_go/lesson28/server/endpoint"
	"github.com/weirubo/intermediate_go/lesson28/server/middleware"
	"github.com/weirubo/intermediate_go/lesson28/server/service"
	"github.com/weirubo/intermediate_go/lesson28/server/transport"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var consulClient *consulApi.Client

func init() {
	config := consulApi.DefaultConfig()
	config.Address = "localhost:8500"
	client, err := consulApi.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	consulClient = client
}

func main() {
	var logger kitLog.Logger
	logger = kitLog.NewLogfmtLogger(os.Stdout)
	logger = kitLog.WithPrefix(logger, "caller", kitLog.DefaultCaller)
	logger = kitLog.WithPrefix(logger, "time", kitLog.DefaultTimestampUTC)

	user := service.User{}
	loginEndpoint := (middleware.LogMiddleware(logger))(endpoint.LoginEndpoint(user))
	handler := transportHttp.NewServer(loginEndpoint, transport.DecodeRequest, transport.EncodeResponse)
	r := mux.NewRouter()
	r.Handle(`/{email}/{password}`, handler)
	r.Methods("GET").Path("/health").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-type", "application/json")
			_, err := w.Write([]byte(`{"status": "ok"}`))
			if err != nil {
				log.Println(err)
				return
			}
		})

	errChan := make(chan error)
	// 服务注册
	go func() {
		ServiceRegister()
		err := http.ListenAndServe(":8080", r)
		if err != nil {
			log.Println(err)
			errChan <- err
		}
	}()

	// 优雅退出
	go func() {
		signalChan := make(chan os.Signal)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-signalChan)
	}()

	// 服务注销
	err := <-errChan
	ServiceDeregister("userServer")
	ServiceDeregister("user")
	log.Println(err)
}

func ServiceRegister() {
	config := consulApi.DefaultConfig()
	config.Address = "localhost:8500"
	registration := consulApi.AgentServiceRegistration{
		Name:    "userServer",
		Address: "localhost",
		Port:    8080,
		Tags:    []string{"userServer"},
		Check: &consulApi.AgentServiceCheck{
			Interval: "5s",
			HTTP:     "http://localhost:8080/health",
		},
	}
	err := consulClient.Agent().ServiceRegister(&registration)
	if err != nil {
		log.Fatal(err)
	}
}

func ServiceDeregister(serviceID string) {
	err := consulClient.Agent().ServiceDeregister(serviceID)
	if err != nil {
		log.Fatal(err)
	}
}
