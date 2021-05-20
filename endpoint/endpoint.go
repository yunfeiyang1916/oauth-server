package endpoint

import "github.com/yunfeiyang1916/oauth-server/service"

// Container 终端容器
// 终端负责接收请求，处理请求，并返回结果。可以添加熔断、日志、限流、负载均衡等能力
type Container struct {
	ClientEndpoint *ClientEndpoint
}

func New(srv *service.Container) *Container {
	return &Container{
		ClientEndpoint: newClientEndpoint(srv.ClientService),
	}
}
