package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

var (
	config *Config
	once   sync.Once
)

type Config struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func InitConfig(FILE_PATH string) {
	once.Do(func() {
		// 导入配置文件
		// fmt.Printf(os.Getwd())
		fmt.Println(FILE_PATH)
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
	})
}

func readConfig(FILE_PATH string) {

}

// Get 提供只读的全局配置
func Get() Config {
	return *config
}
