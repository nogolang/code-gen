package gormUtils

import (
	"code-gen/internal/model"
	"code-gen/internal/utils/commonRes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func GetConnect(model *model.OrmModel) (*gorm.DB, error) {
	config := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}
	url := model.Username + ":" + model.Password + "@tcp(" + model.Host + ":" + model.Port + ")/" + model.DataBaseName + "?charset=utf8mb4&parseTime=True&loc=Local"

	//gormDb无需使用.session，它Open出来就是一个链式安全的实例
	tempDb, err := gorm.Open(mysql.Open(url), config)
	if err != nil {
		return nil, commonRes.DataBaseConnectFail
	}
	return tempDb, nil

}
