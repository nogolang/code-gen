package utils

import "github.com/spf13/viper"

// 如果isProd为True，则是生成环境，返回false
// 如果没有找到isProd，那就是开发环境，返回true
// 如果isProd为false，则是开发环境，返回true
func IsDev() bool {
	err := viper.BindEnv("isProd")
	if err != nil {
		//如果没有绑定，那么当前就是开发环境
		return true
	}

	//查看值，如果为ture，则是生产环境，为false，则是开发环境
	if b := viper.GetBool("isProd"); b {
		return false
	}

	return true
}
