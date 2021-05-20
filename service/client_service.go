package service

import (
	"context"

	"github.com/yunfeiyang1916/oauth-server/api/code"

	"github.com/yunfeiyang1916/toolkit/logging"

	"github.com/yunfeiyang1916/oauth-server/dao"

	"github.com/yunfeiyang1916/oauth-server/model"
)

// ClientService 客户端服务实现
type ClientService struct {
	// 客户端数据访问层
	clientDao *dao.ClientDao
}

func newClientService(clientDao *dao.ClientDao) *ClientService {
	return &ClientService{
		clientDao: clientDao,
	}
}

// Get 通过客户端标识与密钥获取
func (c *ClientService) Get(ctx context.Context, req model.ClientReq) model.BaseResp {
	var (
		resp   = model.BaseResp{ApiResult: model.GetApiResult(code.ErrOK)}
		logger = logging.For(ctx, "func", "get", "req", req)
	)
	if req.Id == "" || req.ClientSecret == "" {
		resp.ApiResult = model.GetApiResult(code.ErrParam)
		return resp
	}
	entity, err := c.clientDao.FindById(req.Id)
	if err != nil {
		logger.Errorf("FindById db error,err=%s", err)
		resp.ApiResult = model.GetApiResult(code.ErrSystem)
		return resp
	}
	resp.Data = entity
	return resp
}
