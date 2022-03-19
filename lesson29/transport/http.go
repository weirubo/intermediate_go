package transport

import (
	"context"
	"encoding/json"
	kitHttp "github.com/go-kit/kit/transport/http"
	"github.com/weirubo/intermediate_go/lesson29/endpoint"
	"net/http"
)

// 对外调用接口(http 或 rpc)

func NewHttpHandler(ctx context.Context, endpoints *endpoint.Endpoints) http.Handler {
	r := http.NewServeMux()
	r.Handle("/register", kitHttp.NewServer(endpoints.RegisterEndpoint, decRegisterRequest, encResponse))
	return r
}

// decode
func decRegisterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	return &endpoint.RegisterRequest{
		UserName: username,
		Email:    email,
		Password: password,
	}, nil
}

// encode
func encResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
