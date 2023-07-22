package service

import (
	"go-start/pkg/model"
	"go-start/pkg/repository"
	"go-start/pkg/service/impl"
)

type DataService interface {
	WriteData(req model.SaveRequest) (int, error)
	GetPropertyByName(name string) (interface{}, error)
}

type Service struct {
	DataService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		DataService: impl.NewDataServiceImpl(repo.DataRepository),
	}
}
