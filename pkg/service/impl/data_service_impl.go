package impl

import (
	"go-start/pkg/model"
	"go-start/pkg/repository"
)

type DataServiceImpl struct {
	repo repository.DataRepository
}

func NewDataServiceImpl(repo repository.DataRepository) *DataServiceImpl {
	return &DataServiceImpl{repo: repo}
}

func (service *DataServiceImpl) WriteData(req model.SaveRequest) (int, error) {
	return service.repo.WriteData(req)
}

func (service *DataServiceImpl) GetPropertyByName(name string) (interface{}, error) {
	return service.repo.GetPropertyByName(name)
}
