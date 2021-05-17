package model

import "time"

// OAuth2Token 令牌
type OAuth2Token struct {
	// 刷新令牌
	RefreshToken *OAuth2Token
	// 令牌类型
	TokenType string
	// 令牌
	TokenValue string
	// 过期时间
	ExpiresTime *time.Time
}

// IsExpired 是否过期
func (o *OAuth2Token) IsExpired() bool {
	return o.ExpiresTime != nil && o.ExpiresTime.Before(time.Now())
}

// OAuth2 认证，两种认证方式，客户端方式以及用户密码方式
type OAuth2 struct {
	// 客户端详情
	Client Client
	// 用户详情
	User User
}
