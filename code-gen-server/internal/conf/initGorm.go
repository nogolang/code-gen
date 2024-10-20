package conf

import (
	"code-gen/internal/model"
	"code-gen/internal/utils/gormUtils"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
)

var GormProvider = wire.NewSet(NewGormConfig)

// NewGormConfig logger由外部注入进来
func NewGormConfig(zapLog *zap.Logger) *gorm.DB {
	type GormConfig struct {
		Url string `json:"url"`
	}

	var receiver GormConfig
	err := viper.UnmarshalKey("gorm", &receiver)
	if err != nil {
		log.Fatal("gorm配置解析失败", err)
		return nil
	}

	//gorm适配zap
	myGormZap := gormUtils.NewMyGormZap(zapLog, gormlogger.Info)

	//初始化gorm的配置
	config := &gorm.Config{
		Logger: myGormZap,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		//禁用迁移的时候自动创建外键
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	//gormDb无需使用.session，它Open出来就是一个链式安全的实例
	db, err := gorm.Open(mysql.Open(receiver.Url), config)
	if err != nil {
		log.Fatal("gorm连接数据库失败", err)
		return nil
	}

	//迁移所有的表
	migrator := db.Migrator()
	err = migrator.AutoMigrate(&model.OutDirModel{},
		&model.MappingPathModel{},
		&model.FileModel{},
		&model.GroupModel{},
		&model.FileAndGroupModel{},
		&model.OrmModel{},
		&model.FileGenModel{},
	)
	if err != nil {
		log.Fatal("表迁移失败", err)
		return nil
	}
	return db
}
