package main

import (
	"code-gen/configs"
	"flag"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "conf", "configs/config.dev.yaml", "config path, eg: -conf configs/config.dev.yaml")
}
func main() {
	//解析配置
	flag.Parse()

	//加载所有配置
	allConfig := configs.ReadConfig(configPath)

	//启动服务
	app := WireApp(allConfig)
	app.RunServer()
}
