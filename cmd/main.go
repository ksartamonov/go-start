package main

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/joho/godotenv"
	_ "github.com/mattes/migrate/source/file"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go-start/pkg/handler"
	"go-start/pkg/repository"
	"go-start/pkg/service"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.TextFormatter))
	logrus.SetLevel(logrus.DebugLevel)

	if err := InitConfig(); err != nil {
		logrus.Fatalf("error while initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading environment variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.DbConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.name"),
	})
	if err != nil {
		logrus.Fatalf("error initializing db: %s", err.Error())
	}

	m, err := migrate.New(viper.GetString("db.migration"), viper.GetString("db.url"))
	if err != nil {
		logrus.Fatalf("error creating migration: %s", err.Error())
	}
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		logrus.Fatalf("up migration error: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos, []string{viper.GetString("kafka.broker")})
	handlers := handler.NewHandler(services)

	server := new(Server)
	err = server.Run(viper.GetString("port"), handlers.RouteHandlers())
	if err != nil {
		logrus.Fatalf("error while running http server: %s", err.Error())
	}
	//
	//if err = m.Down(); err != nil && err != migrate.ErrNoChange {
	//	logrus.Fatalf("down migration error: %s", err.Error())
	//}
}

func InitConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
