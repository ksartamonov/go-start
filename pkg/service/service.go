package service

import (
	"go-start/pkg/model"
	"go-start/pkg/repository"
	"go-start/pkg/service/impl"
)

type DataService interface {
	WriteData(req model.SaveRequest) (int, error)
	GetParameterValue(name string) ([]string, error)
	GetByPair(property model.Property) (model.GetByPairResponse, error)
}

type Service struct {
	DataService
	brokersURL []string // Kafka
}

func NewService(repo *repository.Repository, brokersURL []string) *Service {
	return &Service{
		DataService: impl.NewDataServiceImpl(repo.DataRepository, brokersURL),
		brokersURL:  brokersURL,
	}
}
