package impl

import (
	"go-start/pkg/model"
	"go-start/pkg/repository"
)

type DataServiceImpl struct {
	repo       repository.DataRepository
	brokersURL []string
}

func NewDataServiceImpl(repo repository.DataRepository, brokersURL []string) *DataServiceImpl {
	return &DataServiceImpl{
		repo:       repo,
		brokersURL: brokersURL,
	}
}

func (service *DataServiceImpl) WriteData(req model.SaveRequest) (int, error) {

	result := make(map[string]interface{})
	for _, d := range req.Properties {
		result[d.Parameter] = d.Value
	}

	// Kafka demo
	//reqByte, err := json.Marshal(req)
	//if err != nil {
	//	logrus.Errorf("error marshaling json: %s", err.Error())
	//	return -1, err
	//}
	//err = mq.PushCommentToQueue("comments", reqByte, service.brokersURL)
	//if err != nil {
	//	logrus.Errorf("error writing to queue: %s", err.Error())
	//}

	return service.repo.WriteData(req.Name, result)
}

func (service *DataServiceImpl) GetParameterValue(name string) ([]string, error) {
	return service.repo.GetParameterValue(name)
}

func (service *DataServiceImpl) GetByPair(property model.Property) (model.GetByPairResponse, error) {
	return service.repo.GetByPair(property)
}
