package config

import "fmt"

type Mysql struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"`                               // 数据库主机地址
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`                               // 数据库端口
	User         string `mapstructure:"user" json:"user" yaml:"user"`                               // 数据库用户名
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                   // 数据库密码
	Dbname       string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`                         // 数据库名
	Charset      string `mapstructure:"charset" json:"charset" yaml:"charset"`                      // 字符集，例如 utf8mb4
	MaxOpenConns int    `mapstructure:"max_open_conns" json:"max_open_conns" yaml:"max_open_conns"` // 最大连接数
	MaxIdleConns int    `mapstructure:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns"` // 最大空闲连接数
}

func (m *Mysql) DSN() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" +
		fmt.Sprintf("%d", m.Port) + ")/" + m.Dbname + "?charset=" + m.Charset + "&parseTime=True&loc=Local"
}
