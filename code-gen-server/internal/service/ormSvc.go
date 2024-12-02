package service

import (
	"code-gen/internal/dao"
	"code-gen/internal/model"
	"code-gen/internal/utils/gormUtils"
	"go.uber.org/zap"
)

type OrmSvc struct {
	Logger     *zap.Logger
	Dao        *dao.OrmDao
	FileGenDao *dao.FileGenDao
}

func (receiver *OrmSvc) Add(groupModel *model.OrmModel) error {
	return receiver.Dao.Add(groupModel)
}

func (receiver *OrmSvc) FindById(id int) (*model.OrmModel, error) {
	return receiver.Dao.FindById(id)
}

func (receiver *OrmSvc) DeleteById(id int) error {
	//把fileGen里的dataSources id也置0
	//err := receiver.FileGenDao.SetZeroIdWithDatabaseId(id)
	//if err != nil {
	//	return err
	//}

	return receiver.Dao.DeleteById(id)
}

func (receiver *OrmSvc) UpdateById(id int, m *model.OrmModel) error {
	return receiver.Dao.UpdateById(id, m)
}

func (receiver *OrmSvc) FindAll(query *model.OrmModelQuery) (*model.OrmModelsAll, error) {
	var modelAll model.OrmModelsAll
	data, total, err := receiver.Dao.FindAll(query)
	if err != nil {
		return nil, err
	}

	//如果数量等于0，那么不返回错误，则是返回空数组
	if total == 0 {
		return &modelAll, nil
	}

	modelAll.Data = data
	modelAll.Total = total
	return &modelAll, nil
}

func (receiver *OrmSvc) CheckConnect(m *model.OrmModel) error {
	_, err := gormUtils.GetConnect(m)
	if err != nil {
		return err
	}
	return nil
}

func (receiver *OrmSvc) FindAllNoPagination() ([]model.OrmModel, error) {
	return receiver.Dao.FindAllNoPagination()
}

func (receiver *OrmSvc) FindTablesByDatabaseId(id int) ([]string, error) {
	m, err := receiver.FindById(id)
	if err != nil {
		return nil, err
	}

	//获取到db连接
	db, err := gormUtils.GetConnect(m)
	if err != nil {
		return nil, err
	}
	return receiver.Dao.FindTables(db)
}
