package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/yunfeiyang1916/oauth-server/service"
)

// OAuth2Endpoint 终端，负责接收请求，处理请求，并返回结果。可以添加熔断、日志、限流、负载均衡等能力
type OAuth2Endpoint struct {
	// 令牌终端
	TokenEndpoint endpoint.Endpoint
	// 校验令牌终端
	CheckTokenEndpoint endpoint.Endpoint
}

func newTokenEndpoint(srv *service.Container) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return "ss", nil

	}
}
