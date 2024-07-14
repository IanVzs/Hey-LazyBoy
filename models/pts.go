package models

type PtsParams struct {
	Url         string `json:"url"`
	Count       int    `json:"count"`
	SecondCount int    `json:"secondcount"`
}

type RespPts struct {
	Total       int `json:"total"`
	During      int `json:"during"`
	SecondCount int `json:"secondcount"`
}
