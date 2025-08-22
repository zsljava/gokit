package config

type Server struct {
	Env    string     `mapstructure:"env" json:"env" yaml:"env"`
	Server ServerInfo `mapstructure:"server" json:"server" yaml:"server"`
	JWT    JWT        `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis  Redis      `mapstructure:"redis" json:"redis" yaml:"redis"`
	Mysql  Mysql      `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}
