package config

type Jwt struct {
	Secret        string `mapstructure:"secret" json:"secret" yaml:"secret"`
	AccessExpire  int64  `mapstructure:"access_expire" json:"access_expire" yaml:"access_expire"`
	RefreshExpire int64  `mapstructure:"refresh_expire" json:"refresh_expire" yaml:"refresh_expire"`
}
