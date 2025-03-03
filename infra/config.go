package infra

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

// Config 应用配置文件结构
type Config struct {
	MySQL struct {
		DSN string `yaml:"dsn"` // MySQL 数据源名称
	} `yaml:"mysql"`
	RocketMQ struct {
		NameServer   string `yaml:"name_server"`   // RocketMQ NameServer 地址
		PaymentTopic string `yaml:"payment_topic"` // 支付结果 Topic
		PayoutTopic  string `yaml:"payout_topic"`  // 出款结果 Topic
	} `yaml:"rocketmq"`
}

// NewConfig 读取并解析配置文件
func NewConfig() (Config, error) {
	data, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		return Config{}, err
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}
