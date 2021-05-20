package http

import (
	"context"
	"net/http"

	"github.com/yunfeiyang1916/oauth-server/model"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/yunfeiyang1916/oauth-server/plugins"
)

// getClient 通过客户端标识与密钥获取
func getClient() *kithttp.Server {
	return kithttp.NewServer(plugins.CheckTokenMiddleware()(eptContainer.ClientEndpoint.Get()),
		func(ctx context.Context, r *http.Request) (interface{}, error) {
			// 转换请求
			req := model.ClientReq{}
			req.Id = r.URL.Query().Get("id")
			req.ClientSecret = r.URL.Query().Get("client_secret")
			return req, nil
		},
		encodeJsonResp,
		options...)
}
