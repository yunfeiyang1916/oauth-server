package endpoint

import (
	"context"

	"github.com/yunfeiyang1916/oauth-server/model"

	"github.com/yunfeiyang1916/oauth-server/service"

	"github.com/go-kit/kit/endpoint"
)

// ClientEndpoint 客户端终端
type ClientEndpoint struct {
	service *service.ClientService
}

func newClientEndpoint(service *service.ClientService) *ClientEndpoint {
	return &ClientEndpoint{
		service: service,
	}
}

// Get 通过客户端标识与密钥获取
func (c *ClientEndpoint) Get() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(model.ClientReq)
		return c.service.Get(ctx, req), nil
	}
}
