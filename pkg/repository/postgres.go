package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
)

type DbConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func NewPostgresDB(cfg DbConfig) (*pgx.Conn, error) {
	db, err := pgx.Connect(context.Background(), fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName))
	if err != nil {
		log.Fatalf("error while connecting to the database: %s", err.Error())
	}

	return db, nil
}
