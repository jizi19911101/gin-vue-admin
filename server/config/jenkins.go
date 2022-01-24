package config

type Jenkins struct {
	User  string `mapstructure:"user" json:"user" yaml:"user"`    // Jnekins用户名，用于远程调用jenkins
	Token string `mapstructure:"token" json:"token" yaml:"token"` // Jnekins token，用于远程调用jenkins
}
