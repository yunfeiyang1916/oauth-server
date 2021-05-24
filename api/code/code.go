package code

var e = make(map[int32]string)

const (
	// ErrOK 操作成功
	ErrOK = 200
	// ErrParam 参数错误
	ErrParam = 40000
	// ErrSystem 系统内部错误
	ErrSystem = 50000
)

func init() {
	e[ErrOK] = "操作成功"
	e[ErrParam] = "参数错误"
	e[ErrSystem] = "系统内部错误"
}

// GetErrorMsg 通过code获取错误信息
func GetErrorMsg(code int32) string {
	defMsg := "操作失败，请稍后重试"
	if errMsg, ok := e[code]; ok && errMsg != "" {
		return errMsg
	}
	if ErrOK == code {

	}
	return defMsg
}
