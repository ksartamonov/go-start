package repository

import (
	"github.com/jackc/pgx/v4"
	"go-start/pkg/model"
	"go-start/pkg/repository/impl"
)

type DataRepository interface {
	WriteData(req model.SaveRequest) (int, error)
	GetPropertyByName(name string) (*model.GetResponse, error)
}
type Repository struct {
	DataRepository
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		DataRepository: impl.NewDataRepositoryImpl(db),
	}
}
