package dao

// Container 数据访问层容器
type Container struct {
	// 客户端数据访问层
	ClientDao *ClientDao
}

func New() *Container {
	return &Container{
		ClientDao: newClientDao(),
	}
}
