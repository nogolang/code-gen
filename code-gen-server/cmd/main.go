package main

import (
	"code-gen/cmd/config"
)

func main() {
	//加载所有配置
	config.ReadConfig()

	//启动服务
	app := WireApp()
	app.RunServer()
}
