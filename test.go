package main

import (
	"context"
	"errors"
	"github.com/cucumber/godog"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go-start/pkg/model"
	"go-start/pkg/repository"
	"go-start/pkg/service"
	"os"
)

func main() {
	if err := InitConfig(); err != nil {
		logrus.Fatalf("error while initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading environment variables: %s", err.Error())
	}

	features := "tests/features" // Path to the directory containing the feature files
	writeDataOpts := godog.Options{
		Format: "progress",
		Paths:  []string{features},
	}

	status := godog.TestSuite{
		Name:                "WriteData Endpoint",
		ScenarioInitializer: InitializeScenario,
		Options:             &writeDataOpts,
	}.Run()

	if status != 0 {
		logrus.Fatalf("There are failing scenarios!")
	}

	findByPairOpts := godog.Options{
		Format: "progress",
		Paths:  []string{features},
		Tags:   "@FindByPair",
	}

	findByPairStatus := godog.TestSuite{
		Name:                "FindByPair Endpoint",
		ScenarioInitializer: InitializeScenario,
		Options:             &findByPairOpts,
	}.Run()

	if findByPairStatus != 0 {
		logrus.Fatalf("There are failing scenarios in FindByPair test!")
	}
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^a request to save data$`, aRequestToSaveData)
	ctx.Step(`^the WriteData endpoint is called$`, theWriteDataEndpointIsCalled)
	ctx.Step(`^the data should be written successfully$`, theDataShouldBeWrittenSuccessfully)
	ctx.Step(`^a property with parameter "([^"]*)" and value "([^"]*)"$`, aPropertyWithParameterAndValue)
	ctx.Step(`^the GetByPair endpoint is called$`, theGetByPairEndpointIsCalled)
	ctx.Step(`^the response should contain all entities with the given pair$`, theResponseShouldContainEntitiesWithTheGivenPair)
	ctx.Step(`^the response should have no entities$`, theResponseShouldHaveNoEntities)
	ctx.Step(`^a parameter name "([^"]*)"$`, aParameterName)
	ctx.Step(`^the GetParameterValue endpoint is called$`, theGetParameterValueEndpointIsCalled)
	ctx.Step(`^the response should contain all parameter values$`, theResponseShouldContainAllParameterValues)
	ctx.Step(`^the response should contain no values$`, theResponseShouldContainNoValues)
}

var req = model.SaveRequest{}
var svc *service.Service
var repo *repository.Repository
var result int
var err error

func aRequestToSaveData() error {
	req = model.SaveRequest{
		Name: "Test",
		Properties: []model.Property{
			{Parameter: "param1", Value: "value1"},
			{Parameter: "param2", Value: "value2"},
		},
	}
	return nil
}

func theWriteDataEndpointIsCalled() error {
	db, err := InitDB()
	if err != nil {
		logrus.Errorf("error initializing DB: %s", err.Error())
	}
	repo = repository.NewRepository(db)
	svc = service.NewService(repo, []string{viper.GetString("kafka.broker")})
	result, err = svc.WriteData(req)

	return nil
}

func theDataShouldBeWrittenSuccessfully() error {
	if err != nil && result != -1 {
		return errors.New("data was not written successfully")
	}
	return nil
}

var getByPairProperty model.Property
var getByPairResponse model.GetByPairResponse

func aPropertyWithParameterAndValue(parameter, value string) error {
	getByPairProperty = model.Property{
		Parameter: parameter,
		Value:     value,
	}
	return nil
}

func theGetByPairEndpointIsCalled() error {
	db, err := InitDB()
	if err != nil {
		logrus.Errorf("error initializing DB: %s", err.Error())
	}
	repo := repository.NewRepository(db)
	svc = service.NewService(repo, []string{viper.GetString("kafka.broker")})
	getByPairResponse, err = svc.GetByPair(getByPairProperty)

	return err
}

func theResponseShouldContainEntitiesWithTheGivenPair() error {
	if len(getByPairResponse.Entities) == 0 {
		return errors.New("no entities returned")
	}

	for _, entity := range getByPairResponse.Entities {
		properties := entity.Properties
		if properties[getByPairProperty.Parameter] != getByPairProperty.Value {
			return errors.New("entity does not contain the given pair")
		}
	}

	return nil
}

func theResponseShouldHaveNoEntities() error {
	if len(getByPairResponse.Entities) != 0 {
		return errors.New("entities found, but should not")
	}
	return nil
}

func InitDB() (*pgx.Conn, error) {
	db, err := repository.NewPostgresDB(repository.DbConfig{
		Host:     viper.GetString("testdb.host"),
		Port:     viper.GetString("testdb.port"),
		Username: viper.GetString("testdb.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("testdb.name"),
	})
	if err != nil {
		logrus.Fatalf("error initializing db: %s", err.Error())
	}

	m, err := migrate.New("file://schema", "postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable")
	if err != nil {
		logrus.Fatalf("error creating migration: %s", err.Error())
	}
	if err := m.Down(); err != nil {
		logrus.Fatalf("down migration error: %s", err.Error())
	}

	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		logrus.Fatalf("up migration error: %s", err.Error())
	}

	sqlScript, err := os.ReadFile("schema/prescript.sql")
	if err != nil {
		logrus.Fatalf("error reading SQL script file: %v", err)
	}

	_, err = db.Exec(context.Background(), string(sqlScript))
	if err != nil {
		logrus.Fatalf("error executing SQL script: %v", err)
	}

	return db, err
}

var parameterName string
var parameterValues []string

func aParameterName(paramName string) error {
	parameterName = paramName
	return nil
}

func theGetParameterValueEndpointIsCalled() error {
	db, err := InitDB()
	if err != nil {
		logrus.Errorf("error initializing DB: %s", err.Error())
	}
	repo = repository.NewRepository(db)
	svc = service.NewService(repo, []string{viper.GetString("kafka.broker")})
	parameterValues, err = svc.GetParameterValue(parameterName)
	return nil
}

func theResponseShouldContainAllParameterValues() error {
	if err != nil || len(parameterValues) == 0 {
		return errors.New("failed to retrieve parameter values")
	}
	return nil
}

func theResponseShouldContainNoValues() error {
	if len(parameterValues) != 0 {
		return errors.New("response should contain no paramete values")
	}
	return nil
}

func InitConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
