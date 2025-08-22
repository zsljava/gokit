package config

type ServerInfo struct {
	Name string `mapstructure:"name" json:"name" yaml:"name"` // 应用名称
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
}
