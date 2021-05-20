package service

import "github.com/yunfeiyang1916/oauth-server/model"

// TokenService 令牌服务
type TokenService struct {
}

func newTokenService() *TokenService {
	return &TokenService{}
}

// CreateAccessToken 生成访问令牌
func (t *TokenService) CreateAccessToken(oauth2 *model.OAuth2) (*model.OAuth2Token, error) {
	return nil, nil
}
