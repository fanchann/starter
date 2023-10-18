package config

import "github.com/spf13/viper"

type IConfig interface {
	Get(key string) string
}

type loadConfig struct{}

func (c *loadConfig) Get(key string) string {
	return viper.GetString(key)
}

func NewLoadloadConfig(file ...string) IConfig {
	v := viper.New()
	v.SetConfigFile(file[0])
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return &loadConfig{}
}
