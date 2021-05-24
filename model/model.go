package model

import _code "github.com/yunfeiyang1916/oauth-server/api/code"

// ApiResult 请求结果
type ApiResult struct {
	// 错误码
	Code int32 `json:"code"`
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
func GetApiResult(code int32) ApiResult {
	msg := _code.GetErrorMsg(code)
	if msg == "" {
		msg = "操作失败，请稍后重试"
	}
	return ApiResult{Code: code, Msg: msg}
}

// GetApiResultWithMsg 获取Api响应接口,包括错误提示语
func GetApiResultWithMsg(code int32, msg string) ApiResult {
	return ApiResult{Code: code, Msg: msg}
}

// GetResp 获取响应基类
func GetResp(code int32, data interface{}) BaseResp {
	return BaseResp{ApiResult: GetApiResult(code), Data: data}
}

// GetRespWithMsg 获取响应基类,包括错误提示语
func GetRespWithMsg(code int32, msg string, data interface{}) BaseResp {
	return BaseResp{ApiResult: GetApiResultWithMsg(code, msg), Data: data}
}
