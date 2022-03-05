package transport

import (
	"context"
	"encoding/json"
	"github.com/weirubo/intermediate_go/lesson28/client/endpoint"
	"net/http"
)

func Req(ctx context.Context, request *http.Request, r interface{}) error {
	req := r.(endpoint.Request)
	request.URL.Path += "/" + req.Email + "/" + req.Password
	return nil
}

func Res(ctx context.Context, response *http.Response) (r interface{}, err error) {
	var res endpoint.Response
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
