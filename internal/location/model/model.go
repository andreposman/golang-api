package model

type Location struct {
	OrderID string `json:"order_id"`
	History []History `json:"history"`
}

type History struct {
	Lat float64 `json:"lat""`
	Lng float64 `json:"lng"`
}
