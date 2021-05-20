package conf

import (
	"github.com/yunfeiyang1916/toolkit/logging"
	"github.com/yunfeiyang1916/toolkit/tomlconfig"
)

// GlobalConfig 全局配置
var GlobalConfig *Config

// Config 配置
type Config struct {
	// 运行环境
	Env string `toml:"env"`
	// 服务配置
	ServerConfig *ServerConfig `toml:"server"`
	LogConfig    *LogConfig    `toml:"log"`
	// 数据库配置
	DBConfig []DBConfig `toml:"database"`
}

// ServerConfig 服务配置
type ServerConfig struct {
	// 服务名
	ServiceName string `toml:"service_name"`
	// 端口
	Port int `toml:"port"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level string `toml:"level"`
	// 配置目录
	LogPath string `toml:"logpath"`
	Rotate  string `toml:"rotate"`
}

// DBConfig 数据库配置
type DBConfig struct {
	// 名称
	Name string `toml:"name"`
}

func Init(path string) (*Config, error) {
	cfg := &Config{}
	_, err := tomlconfig.NewConfig(path, &cfg)
	GlobalConfig = cfg
	initLogging(cfg.LogConfig)
	return cfg, err
}

// 初始化日志
func initLogging(logConfig *LogConfig) {
	logging.SetLevelByString(logConfig.Level)
	logging.SetOutputPath(logConfig.LogPath)
	logging.SetRotateByHour()
}
