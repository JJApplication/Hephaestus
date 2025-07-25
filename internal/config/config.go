package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	AppRoot    string `yaml:"app_root" mapstructure:"app_root"`
	DeployPath string `yaml:"deploy_path" mapstructure:"deploy_path"`
}

var conf *Config

func LoadConfig() *Config {
	if conf != nil {
		return conf
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")

	viper.SetDefault("app_root", "app")
	viper.SetDefault("deploy_path", "deploy")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		// 记录日志 无法获取配置 使用默认配置
	}

	_ = viper.Unmarshal(&conf)

	if conf == nil {
		return &Config{}
	}

	return conf
}
