package impl

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

type DataRepositoryImpl struct {
	db *pgx.Conn
}

func NewDataRepositoryImpl(db *pgx.Conn) *DataRepositoryImpl {
	return &DataRepositoryImpl{db: db}
}

// WriteData returns only id of inserted property (no info about parameters)
func (repo *DataRepositoryImpl) WriteData(name string, parameters map[string]interface{}) (int, error) {
	var id int // id of inserted property

	query, _, err := sq.Insert("data").
		Columns("name", "parameters").
		Values(name, parameters).Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		logrus.Errorf("error building INSERT query: %s", err.Error())
	}
	row := repo.db.QueryRow(context.Background(), query, name, parameters)

	if err := row.Scan(&id); err != nil {
		logrus.Errorf("error getting id from row: %s", err.Error())
		return -1, err
	}

	return id, nil
}

func (repo *DataRepositoryImpl) GetParameterValue(parameterName string) ([]string, error) {
	//query := `SELECT parameters->>$1 FROM data WHERE parameters ? $1 LIMIT 1`
	query := `SELECT parameters->>$1 FROM data WHERE parameters ? $1`

	rows, err := repo.db.Query(context.Background(), query, parameterName)
	if err != nil {
		logrus.Errorf("error executing SQL query: %s", err.Error())
		return nil, err
	}

	var result []string

	for rows.Next() {
		var value string
		err := rows.Scan(&value)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		result = append(result, value)
	}
	return result, nil
}
