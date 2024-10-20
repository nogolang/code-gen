package utils

import "github.com/spf13/viper"

// 如果isDev为True，则是开发环境，返回true
// 如果没有找到isDev，那就是开发环境，返回true
// 如果isDev为false，则是生产环境，返回false
func IsDev() bool {
	//绑定环境变量,我们要在机器上设置了一个isDev的环境变量
	//注意，需要重新打开idea
	//如果绑定失败，那就是没设置，那就是开发环境
	err := viper.BindEnv("isDev")
	if err != nil {
		//如果没有绑定，那么当前就是开发环境
		return true
	}

	//如果绑定成功，则取里面的数值，看看设置的是什么
	if b := viper.GetBool("isDev"); b {
		return true
	}

	return false
}
