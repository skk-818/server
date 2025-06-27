package config

type Logger struct {
	Level        string `mapstructure:"level" json:"level" yaml:"level"` // debug/info/warn/error
	Prefix       string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Format       string `mapstructure:"format" json:"format" yaml:"format"`       // console/json
	Director     string `mapstructure:"director" json:"director" yaml:"director"` // 日志目录
	ShowLine     bool   `mapstructure:"show_line" json:"show_line" yaml:"show_line"`
	LogInConsole bool   `mapstructure:"log_in_console" json:"log_in_console" yaml:"log_in_console"`
}
