package model

type GetByPairResponse struct {
	Entities []Entity `json:"entities"`
}

type Entity struct {
	Id         int               `json:"id"`
	Name       string            `json:"name"`
	Properties map[string]string `json:"properties"`
}
