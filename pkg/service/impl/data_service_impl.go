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

	result := make(map[string]interface{})
	for _, d := range req.Properties {
		result[d.Parameter] = d.Value
	}

	return service.repo.WriteData(req.Name, result)
}

func (service *DataServiceImpl) GetParameterValue(name string) ([]string, error) {
	return service.repo.GetParameterValue(name)
}
