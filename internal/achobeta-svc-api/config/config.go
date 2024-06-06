package config

import (
	"github.com/spf13/viper"
)

var (
	config *Config
	// once   sync.Once
)

type Config struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func InitConfig(FILE_PATH string) {
	// 导入配置文件
	config = &Config{}
	viper.SetConfigFile(FILE_PATH)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
	// 将配置文件读取到结构体中
	err = viper.Unmarshal(config)
	if err != nil {
		panic(err.Error())
	}
}

// Get 提供只读的全局配置
func Get() Config {
	return *config
}
