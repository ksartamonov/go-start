package testcontainers

//
//import (
//	"context"
//	"embed"
//	"errors"
//	"fmt"
//	"github.com/golang-migrate/migrate/v4"
//	_ "github.com/golang-migrate/migrate/v4/database/pgx"
//	_ "github.com/golang-migrate/migrate/v4/database/postgres"
//	_ "github.com/golang-migrate/migrate/v4/source/file"
//	"github.com/golang-migrate/migrate/v4/source/iofs"
//	"github.com/jackc/pgx/v4"
//	"github.com/sirupsen/logrus"
//	"github.com/testcontainers/testcontainers-go"
//	"github.com/testcontainers/testcontainers-go/wait"
//	"strings"
//	"time"
//)
//
//func SetupTestDatabase() (testcontainers.Container, *pgx.Conn) {
//	containerReq := testcontainers.ContainerRequest{
//		Image:        "postgres:latest",
//		ExposedPorts: []string{"5432/tcp"},
//		WaitingFor:   wait.ForLog("database system is ready to accept connections"),
//		Env: map[string]string{
//			"POSTGRES_DB":       "testdb",
//			"POSTGRES_PASSWORD": "postgres",
//			"POSTGRES_USER":     "postgres",
//		},
//	}
//
//	logrus.Debug("Starting PostgreSQL container")
//	dbContainer, err := testcontainers.GenericContainer(
//		context.Background(),
//		testcontainers.GenericContainerRequest{
//			ContainerRequest: containerReq,
//			Started:          true,
//		})
//	if err != nil {
//		logrus.Errorf("error starting PostgreSQL container: %s", err.Error())
//	}
//
//	logrus.Debug("Getting host and port of PostgreSQL container")
//	host, err := dbContainer.Host(context.Background())
//	if err != nil {
//		logrus.Errorf("error getting host of PostgreSQL container: %s", err)
//	}
//	port, err := dbContainer.MappedPort(context.Background(), "5432")
//	if err != nil {
//		logrus.Errorf("error getting port of PostgreSQL container: %s", err)
//	}
//
//	dbURI := fmt.Sprintf("postgres://postgres:postgres@%v:%v/testdb", host, port.Port())
//
//	if err = MigrateDb(dbURI); err != nil {
//		logrus.Errorf("error migrating db: %s", err.Error())
//	}
//
//	db, err := pgx.Connect(context.Background(), dbURI)
//
//	if err != nil {
//		logrus.Errorf("error connecting to DB: %s", err.Error())
//	}
//	return dbContainer, db
//}
//
////go:embed migrations
//var migrations embed.FS
//
//func MigrateDb(dbURI string) error {
//	source, err := iofs.New(migrations, "migrations")
//	if err != nil {
//		logrus.Errorf("error getting driver: %s", err.Error())
//	}
//	//m, err := migrate.NewWithSourceInstance("iofs", source, strings.Replace(dbURI, "postgres://", "pgx://", 1))
//	time.Sleep(15 * time.Second)
//	m, err := migrate.NewWithSourceInstance("iofs", source, strings.Replace(dbURI, "postgres://", "pgx://", 1))
//	if err != nil {
//		logrus.Errorf("error creating Migration instance: %s", err.Error())
//	}
//	defer m.Close()
//
//	err = m.Up()
//	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
//		return err
//	}
//
//	return nil
//}
