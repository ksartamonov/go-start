package repository

import (
	"github.com/jackc/pgx/v4"
	"go-start/pkg/model"
	"go-start/pkg/repository/impl"
)

type DataRepository interface {
	WriteData(name string, parameters map[string]interface{}) (int, error)
	GetParameterValue(parameterName string) ([]string, error)
	GetByPair(property model.Property) (model.GetByPairResponse, error)
}
type Repository struct {
	DataRepository
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		DataRepository: impl.NewDataRepositoryImpl(db),
	}
}
