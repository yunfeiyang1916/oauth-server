package plugins

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// CheckTokenMiddleware 校验令牌中间件
func CheckTokenMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			// return model.GetApiResult(code.ErrParam), nil
			return next(ctx, request)
		}
	}
}

// CheckAuthMiddleware 校验权限中间件
func CheckAuthMiddleware(authority string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			return next(ctx, request)
		}
	}
}
