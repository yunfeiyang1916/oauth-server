package dao

import (
	"fmt"

	"github.com/yunfeiyang1916/oauth-server/model"
	"github.com/yunfeiyang1916/toolkit/sql"
)

// ClientDao 客户端数据访问层
type ClientDao struct {
}

func newClientDao() *ClientDao {
	return &ClientDao{}
}

// FindById 根据id获取
func (c *ClientDao) FindById(id string) (model.Client, error) {
	db := sql.Get("oauth").Master()
	fmt.Println(db)
	return model.Client{}, nil
}
