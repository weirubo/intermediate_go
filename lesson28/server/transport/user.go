package transport

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/weirubo/intermediate_go/lesson28/server/endpoint"
	"net/http"
)

func DecodeRequest(c context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	return endpoint.Request{
		Email:    vars["email"],
		Password: vars["password"],
	}, nil
}

func EncodeResponse(c context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
