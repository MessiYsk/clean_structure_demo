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
	SQLite struct {
		Path string `yaml:"path"` // SQLite 数据库文件路径
	} `yaml:"sqlite"`
}

// NewConfig 读取并解析配置文件
func NewConfig() Config {
	data, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		panic(err)
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}
	return cfg
}
