package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Logger *Logger     `mapstructure:"logger" json:"logger" yaml:"logger"`
	MySQL  *Mysql      `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Server *HTTPServer `mapstructure:"server" json:"server" yaml:"server"`
}

const FilePath = "./etc/config.yaml" // 配置文件路径常量

func LoadConfig() (*Config, error) {
	v := viper.New()

	// 直接指定配置文件路径和名称
	v.SetConfigFile(FilePath) // 这里用常量

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config file error: %w", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config error: %w", err)
	}

	return &cfg, nil
}

// ProvideLoggerConfig 拆解 logger
func ProvideLoggerConfig(cfg *Config) *Logger {
	return cfg.Logger
}

// ProvideMysqlConfig 拆解 logger
func ProvideMysqlConfig(cfg *Config) *Mysql {
	return cfg.MySQL
}

func ProvideHttpServerConfig(cfg *Config) *HTTPServer {
	return cfg.Server
}
