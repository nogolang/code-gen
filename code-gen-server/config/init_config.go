package config

import (
	"code-gen/internal/utils"
	"github.com/spf13/viper"
	"log"
	"os"
)

// ReadConfig 读取所有配置
func ReadConfig() {
	isExists := utils.IsDev()
	dir, _ := os.Getwd()
	log.Println("当前路径：", dir)
	if isExists {
		log.Println("====当前是开发环境====")
		viper.SetConfigFile(dir + "/config/config.dev.yaml")
	} else {
		log.Println("====当前是生产环境====")
		viper.SetConfigFile(dir + "/config/config.prod.yaml")
	}

	//读取所有配置
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("配置文件格式不正确:", err)
	}
}
