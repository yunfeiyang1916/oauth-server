package dao

import "github.com/yunfeiyang1916/oauth-server/model"

// ClientDao 客户端数据访问层
type ClientDao struct {
}

func newClientDao() *ClientDao {
	return &ClientDao{}
}

// FindById 根据id获取
func (c *ClientDao) FindById(id string) (model.Client, error) {
	return model.Client{}, nil
}
