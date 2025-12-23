package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Logger      *Logger     `mapstructure:"logger" json:"logger" yaml:"logger"`
	SystemMySQL *Mysql      `mapstructure:"system_mysql" json:"system_mysql" yaml:"system_mysql"`
	ImMySQL     *Mysql      `mapstructure:"im_mysql" json:"im_mysql" yaml:"im_mysql"`
	Http        *HTTPServer `mapstructure:"http" json:"http" yaml:"http"`
	Jwt         *Jwt        `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis       *Redis      `mapstructure:"redis" json:"redis" yaml:"redis"`
}

func LoadConfig(path string) (*Config, error) {
	v := viper.New()

	// 直接指定配置文件路径和名称
	v.SetConfigFile(path) // 这里用常量

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config file error: %w", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config error: %w", err)
	}

	return &cfg, nil
}

func ProvideLoggerConfig(cfg *Config) *Logger {
	return cfg.Logger
}

func ProvideMysqlConfig(cfg *Config) *Mysql {
	return cfg.SystemMySQL
}

func ProvideHttpServerConfig(cfg *Config) *HTTPServer {
	return cfg.Http
}

func ProviderCorsConfig(cfg *Config) *Cors {
	return cfg.Http.Cors
}

func ProvideJwtConfig(cfg *Config) *Jwt {
	return cfg.Jwt
}

func ProvideRedisConfig(cfg *Config) *Redis {
	return cfg.Redis
}

func ProvideImMysqlConfig(cfg *Config) *Mysql {
	return cfg.ImMySQL
}
