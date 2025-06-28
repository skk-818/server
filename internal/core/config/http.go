package config

type HTTPServer struct {
	Addr         string `mapstructure:"addr" json:"addr" yaml:"addr"`                            // 服务监听地址
	ReadTimeout  int    `mapstructure:"read_timeout" json:"read_timeout" yaml:"read_timeout"`    // 读取超时时间（秒）
	WriteTimeout int    `mapstructure:"write_timeout" json:"write_timeout" yaml:"write_timeout"` // 写入超时时间（秒）
}
