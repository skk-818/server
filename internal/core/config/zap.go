package config

type Logger struct {
	Level        string `mapstructure:"level" json:"level" yaml:"level"` // debug/info/warn/error
	Prefix       string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Format       string `mapstructure:"format" json:"format" yaml:"format"`       // console/json
	Director     string `mapstructure:"director" json:"director" yaml:"director"` // 日志目录
	ShowLine     bool   `mapstructure:"show_line" json:"show_line" yaml:"show_line"`
	LogInConsole bool   `mapstructure:"log_in_console" json:"log_in_console" yaml:"log_in_console"`
	MaxSize      int    `mapstructure:"max_size" json:"max_size" yaml:"max_size"`          // 单个日志文件最大MB
	MaxBackups   int    `mapstructure:"max_backups" json:"max_backups" yaml:"max_backups"` // 保留备份数
	MaxAge       int    `mapstructure:"max_age" json:"max_age" yaml:"max_age"`             // 保留天数
	Compress     bool   `mapstructure:"compress" json:"compress" yaml:"compress"`          // 是否压缩
}
