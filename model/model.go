package model

import "github.com/yunfeiyang1916/oauth-server/define"

// ApiResult 请求结果
type ApiResult struct {
	// 错误码
	Code int `json:"code"`
	// 错误提示语
	Msg string `json:"msg"`
}

// BaseResp 带数据的响应基类
type BaseResp struct {
	ApiResult
	Data interface{} `json:"data"`
}

// Page 分页
type Page struct {
	// 页码
	PageIndex int `json:"page_index"`
	// 每页显示数量
	PageSize int `json:"page_size"`
}

// GetApiResult 获取Api响应接口
func GetApiResult(code int) ApiResult {
	msg := define.GetErrorMsg(code)
	if msg == "" {
		msg = "操作失败，请稍后重试"
	}
	return ApiResult{Code: code, Msg: msg}
}

// GetApiResultWithMsg 获取Api响应接口,包括错误提示语
func GetApiResultWithMsg(code int, msg string) ApiResult {
	return ApiResult{Code: code, Msg: msg}
}

// GetResp 获取响应基类
func GetResp(code int, data interface{}) BaseResp {
	return BaseResp{ApiResult: GetApiResult(code), Data: data}
}

// GetRespWithMsg 获取响应基类,包括错误提示语
func GetRespWithMsg(code int, msg string, data interface{}) BaseResp {
	return BaseResp{ApiResult: GetApiResultWithMsg(code, msg), Data: data}
}
