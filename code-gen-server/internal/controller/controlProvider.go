package controller

import (
	"code-gen/internal/dao"
	"code-gen/internal/service"
	"github.com/google/wire"
)

// 后续创建新的controller也需要写在这里
var ProviderSet = wire.NewSet(FileProvider,
	GroupProvider,
	OrmProvider,
	MappingProvider,
	FileGenProvider,
)

// engine由外部注入进来即可
var FileProvider = wire.NewSet(
	NewFileController,
	//接口必须绑定实现类，我们才能使用接口低层模块的接口
	wire.Struct(new(service.FileService), "*"),
	wire.Struct(new(dao.FileDao), "*"),
	wire.Struct(new(dao.MappingPathDao), "*"),
	wire.Struct(new(dao.OutDir), "*"),
)

var GroupProvider = wire.NewSet(
	NewGroupController,
	//接口必须绑定实现类，我们才能使用接口低层模块的接口
	wire.Struct(new(service.GroupSvc), "*"),
	wire.Struct(new(dao.GroupDao), "*"),
	wire.Struct(new(dao.FileAndGroupDao), "*"),
)

var OrmProvider = wire.NewSet(
	NewOrmController,
	wire.Struct(new(service.OrmSvc), "*"),
	wire.Struct(new(dao.OrmDao), "*"),
)

var MappingProvider = wire.NewSet(
	NewMappingController,
	wire.Struct(new(service.MappingSvc), "*"),
)
var FileGenProvider = wire.NewSet(
	NewFileGenController,
	wire.Struct(new(service.FileGenSvc), "*"),
	wire.Struct(new(dao.FileGenDao), "*"),
)
