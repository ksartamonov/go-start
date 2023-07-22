package model

type SaveRequest struct {
	Name       string     `json:"name"`
	Properties []Property `json:"properties"`
}

type Property struct {
	Parameter string `json:"parameter"`
	Value     string `json:"value"`
}
