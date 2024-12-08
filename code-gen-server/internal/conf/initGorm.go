package conf

import (
	"code-gen/configs"
	"code-gen/internal/model"
	"code-gen/internal/utils/gormUtils"
	"github.com/google/wire"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
)

var GormProvider = wire.NewSet(NewGormConfig)

// NewGormConfig logger由外部注入进来
func NewGormConfig(allConfig *configs.AllConfig, logger *zap.Logger) *gorm.DB {

	//gorm适配zap，这里的日志级别还是要取决于我们的是zap
	//因为zap是最终输出的，所以你在这里设置日志级别是无效的
	myGormZap := gormUtils.NewMyGormZap(logger, gormlogger.Info)

	//初始化gorm的配置
	config := &gorm.Config{
		Logger: myGormZap,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		//不自动创建外键
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	//gormDb无需使用.session，它Open出来就是一个链式安全的实例
	db, err := gorm.Open(mysql.Open(allConfig.Gorm.Url), config)
	if err != nil {
		log.Fatal("gorm连接数据库失败", err)
		return nil
	}

	//迁移一些表
	migrator := db.Migrator()
	err = migrator.AutoMigrate(
		&model.MappingPathModel{},
		&model.FileModel{},
		&model.GroupModel{},
		&model.OrmModel{},
		&model.FileGenModel{},
	)
	logger.Info("连接mysql成功")
	return db
}
