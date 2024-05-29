package config

import (
	"github.com/spf13/viper"
)

var (
	config *Config
	// once   sync.Once
)

type Config struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Database struct {
		Mysql struct {
			Host     string `yaml:"host"`
			Port     int    `yaml:"port"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
			Database string `yaml:"database"`
		} `yaml:"mysql"`
		Redis struct {
			Host     string `yaml:"host"`
			Port     int    `yaml:"port"`
			Password string `yaml:"password"`
			Database int    `yaml:"database"`
		} `yaml:"redis"`
	} `yaml:"database"`
}

func InitConfig(FILE_PATH string) {
	// once.Do(func() {
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
	// fmt.Printf("config: %+v", config)
	// })
}

// Get 提供只读的全局配置
func Get() Config {
	return *config
}
