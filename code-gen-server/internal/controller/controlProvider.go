package controller

import (
	"code-gen/internal/dao"
	"code-gen/internal/service"
	"github.com/google/wire"
)

// 后续创建新的controller也需要写在这里
var ProviderSet = wire.NewSet(
	GroupProvider,
	OrmProvider,
	MappingProvider,
	FileGenProvider,
	StaticProvider,
)

// 静态文件服务
var StaticProvider = wire.NewSet()

var GroupProvider = wire.NewSet(
	NewGroupController,
	//接口必须绑定实现类，我们才能使用接口低层模块的接口
	wire.Struct(new(service.GroupSvc), "*"),
	wire.Struct(new(dao.GroupDao), "*"),
	wire.Struct(new(dao.FileDao), "*"),
)

var OrmProvider = wire.NewSet(
	NewOrmController,
	wire.Struct(new(service.OrmSvc), "*"),
	wire.Struct(new(dao.OrmDao), "*"),
)

var MappingProvider = wire.NewSet(
	NewMappingController,
	wire.Struct(new(service.MappingSvc), "*"),
	wire.Struct(new(dao.MappingPathDao), "*"),
)
var FileGenProvider = wire.NewSet(
	NewFileGenController,
	wire.Struct(new(service.FileGenSvc), "*"),
	wire.Struct(new(dao.FileGenDao), "*"),
)
