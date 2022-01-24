package config

type Gitlab struct {
	ChumanApiTestUrl string `mapstructure:"chuman-api-test-url" json:"chumanApiTestUrl" yaml:"chuman-api-test-url"` // gitlab触漫接口自动化代码地址
}
