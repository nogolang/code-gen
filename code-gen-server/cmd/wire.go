//go:build wireinject
// +build wireinject

package main

import (
	"code-gen/internal/conf"
	"code-gen/internal/controller"
	"github.com/google/wire"
)

// 注入http服务
func WireApp() *conf.HttpServer {
	wire.Build(
		wire.Struct(new(conf.HttpServer), "*"),
		conf.ZapProvider,
		conf.GormProvider,
		conf.NewGin,
		//把所有controller汇集到一起
		controller.ProviderSet,
	)
	return &conf.HttpServer{}
}
