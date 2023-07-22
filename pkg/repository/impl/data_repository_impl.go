package impl

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	"go-start/pkg/model"
)

type DataRepositoryImpl struct {
	db *pgx.Conn
}

func NewDataRepositoryImpl(db *pgx.Conn) *DataRepositoryImpl {
	return &DataRepositoryImpl{db: db}
}

// WriteData returns only id of inserted property (no info about parameters)
func (repo *DataRepositoryImpl) WriteData(req model.SaveRequest) (int, error) {
	var propertyId int // id of inserted property
	query := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", "property")
	row := repo.db.QueryRow(context.Background(), query, req.Name)

	if err := row.Scan(&propertyId); err != nil {
		logrus.Errorf("error getting id from row: %s", err.Error())
		return -1, err
	}

	// inserting parameters
	query = fmt.Sprintf("INSERT INTO %s (parameter, value, property_id) VALUES ($1, $2, $3)", "parameter")
	for i := 0; i < len(req.Properties); i++ { // TODO: think may be optimized
		_, err := repo.db.Exec(context.Background(), query, req.Properties[i].Parameter, req.Properties[i].Value, propertyId)
		if err != nil {
			logrus.Errorf("error writing to parameter table: %s", err.Error())
			return -1, err
		}
	}

	return propertyId, nil
}

func (repo *DataRepositoryImpl) GetPropertyByName(name string) (*model.GetResponse, error) {
	query := `
		SELECT p.id, p.name, pa.parameter, pa.value
		FROM property p
		LEFT JOIN parameter pa ON p.id = pa.property_id
		WHERE p.name = $1
	`
	rows, err := repo.db.Query(context.Background(), query, name)
	if err != nil {
		logrus.Errorf("error executing query: %s", err.Error())
		return nil, err
	}
	defer rows.Close()

	getResp := &model.GetResponse{}
	parameters := make([]model.Property, 0)

	for rows.Next() {
		var id int
		var name string
		var parameter string
		var value string

		err := rows.Scan(&id, &name, &parameter, &value)
		if err != nil {
			logrus.Errorf("error scanning row: %s", err.Error())
			return nil, err
		}

		getResp.Id = id
		getResp.Name = name
		parameters = append(parameters, model.Property{
			Parameter: parameter,
			Value:     value,
		})
	}

	if err := rows.Err(); err != nil {
		logrus.Errorf("error iterating over rows: %s", err.Error())
		return nil, err
	}

	getResp.Properties = parameters

	return getResp, nil
}
