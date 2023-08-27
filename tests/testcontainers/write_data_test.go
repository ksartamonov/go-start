package testcontainers

//
//import (
//	"context"
//	"github.com/sirupsen/logrus"
//	"go-start/pkg/model"
//	"go-start/pkg/repository"
//	"go-start/pkg/service"
//	"testing"
//)
//
//var properties = []model.Property{
//	{Parameter: "test_parameter_1", Value: "test_data_1"},
//	{Parameter: "test_parameter_2", Value: "test_data_2"},
//}
//
//var req = model.SaveRequest{
//	Name:       "test",
//	Properties: properties,
//}
//
//func TestDataServiceImpl(t *testing.T) {
//	logrus.SetLevel(logrus.DebugLevel)
//	logrus.Debug("Creating db container")
//	dbContainer, db := SetupTestDatabase()
//
//	defer dbContainer.Terminate(context.Background())
//
//	repo := repository.NewRepository(db)
//	serv := service.NewService(repo, nil)
//
//	t.Run("WriteData", func(t *testing.T) {
//		data, _ := serv.WriteData(req)
//		logrus.Debug("data = ", data)
//		logrus.Debugf("hihihih")
//		// assert.NotEqualf(t, -1, data, "error, id cannot be -1")
//	})
//}
