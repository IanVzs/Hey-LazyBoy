package models

type PtsParams struct {
	Url         string `json:"url"`
	Count       int    `json:"count"`
	SecondCount int    `json:"secondcount"`
	Duration    int    `json:"duration"`
}

type RespPts struct {
	Total       int `json:"total"`
	Cost        int `json:"cost"`
	SecondCount int `json:"secondcount"`
}
