package model

// Client 客户端
type Client struct {
	// 客户端标识
	Id string
	// 客户端的密钥
	ClientSecret string
	// 访问令牌有效时间，秒
	AccessTokenValiditySeconds int
	// 刷新令牌有效时间，秒
	RefreshTokenValiditySeconds int
	// 重定向地址，授权码类型中使用
	RegisteredRedirectUri string
	// 可以使用的授权类型
	AuthorizedGrantTypes []string
}

// ClientReq 请求
type ClientReq struct {
	// 客户端标识
	Id string `json:"id"`
	// 客户端的密钥
	ClientSecret string `json:"client_secret"`
}

// ClientResp 客户端响应
type ClientResp struct {
	ApiResult
	Data Client `json:"data"`
}
