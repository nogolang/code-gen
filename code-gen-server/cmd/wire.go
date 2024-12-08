//go:build wireinject
// +build wireinject

package main

import (
	"code-gen/configs"
	"code-gen/internal/conf"
	"code-gen/internal/controller"
	"github.com/google/wire"
)

// 注入http服务
func WireApp(allConfig *configs.AllConfig) *conf.HttpServer {
	wire.Build(
		wire.Struct(new(conf.HttpServer), "*"),
		//所有配置都汇集到一起，包括gin对象的创建
		conf.ProviderSet,

		//所有controller汇集到一起
		controller.ProviderSet,
	)
	return &conf.HttpServer{}
}
