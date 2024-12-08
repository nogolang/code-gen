package service

import (
	"code-gen/internal/dao"
	"code-gen/internal/model"
	"go.uber.org/zap"
)

type MappingSvc struct {
	Logger  *zap.Logger
	Dao     *dao.MappingPathDao
	FileDao *dao.FileDao
}

func (receiver *MappingSvc) Add(m *model.MappingPathModel) error {
	return receiver.Dao.Add(m)
}

func (receiver *MappingSvc) FindById(id int) (*model.MappingPathModel, error) {
	return receiver.Dao.FindById(id)
}

func (receiver *MappingSvc) DeleteById(id int) error {
	return receiver.Dao.DeleteById(id)
}

func (receiver *MappingSvc) UpdateById(id int, m *model.MappingPathModel) error {
	return receiver.Dao.UpdateById(id, m)
}

func (receiver *MappingSvc) FindAll(query *model.MappingPathModelQuery) (*model.MappingPathModelsAll, error) {
	var modelAll model.MappingPathModelsAll
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

func (receiver *MappingSvc) FindAllNoPagination() ([]model.MappingPathModel, error) {
	return receiver.Dao.FindAllNoPagination()

}
