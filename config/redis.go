package config

type Redis struct {
	Addr         string `mapstructure:"addr" json:"addr" yaml:"addr"`                            // 服务器地址:端口
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                // 密码
	DB           int    `mapstructure:"db" json:"db" yaml:"db"`                                  // 单实例模式下redis的哪个数据库
	Readtimeout  string `mapstructure:"read_timeout" json:"read_timeout" yaml:"read_timeout"`    // 读超时时间
	WriteTimeout string `mapstructure:"write_timeout" json:"write_timeout" yaml:"write_timeout"` // 写超时时间
}
