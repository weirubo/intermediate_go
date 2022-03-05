package middleware

import (
	"context"
	kitEndpoint "github.com/go-kit/kit/endpoint"
	kitLog "github.com/go-kit/kit/log"
	"github.com/weirubo/intermediate_go/lesson28/server/endpoint"
)

func LogMiddleware(logger kitLog.Logger) kitEndpoint.Middleware {
	return func(after kitEndpoint.Endpoint) kitEndpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			r := request.(endpoint.Request)
			logger.Log("email", r.Email)
			return after(ctx, request)
		}
	}
}
