package model

// User 用户
type User struct {
	Id int64
	// 用户名 唯一
	Username string
	// 用户密码
	Password string
	// 用户具有的权限
	Authorities []string
}

