package configs

import (
	"github.com/spf13/viper"
	"log"
)

// ReadConfig 读取所有配置
func ReadConfig(configPath string) *AllConfig {
	viper.SetConfigFile(configPath)

	//读取所有配置
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("配置文件格式不正确:", err)
	}

	var allConfig AllConfig
	err = viper.Unmarshal(&allConfig)
	if err != nil {
		log.Fatal("配置文件解析失败:", err)
	}
	return &allConfig
}
