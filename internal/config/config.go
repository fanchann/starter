package config

import "github.com/spf13/viper"

// var (
// 	v *viper.Viper
// )

type IConfig interface {
	Get(key string) string
}

type loadConfig struct{}

func (c *loadConfig) Get(key string) string {
	return viper.GetString(key)
}

func NewLoadConfig(file ...string) IConfig {
	viper.SetConfigFile(file[0])

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return &loadConfig{}
}
