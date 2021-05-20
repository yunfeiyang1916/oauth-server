package conf

// Config 配置
type Config struct {
}

func Init() (*Config, error) {
	cfg := &Config{}
	// 从配置文件读取
	return cfg, nil
}
