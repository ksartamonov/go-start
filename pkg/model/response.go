package model

type GetResponse struct {
	Id         int        `json:"id"`
	Name       string     `json:"name"`
	Properties []Property `json:"properties"`
}
