package service

import (
	"github.com/yunfeiyang1916/oauth-server/dao"
)

// Container 服务容器
type Container struct {
	// 客户端服务
	ClientService *ClientService
}

func New() *Container {
	d := dao.New()
	return &Container{
		ClientService: newClientService(d.ClientDao),
	}
}

// Close close the resource
func (s *Container) Close() {

}
