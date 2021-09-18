package model

type Location struct {
	OrderID string `json:"order_id"`
	History []History `json:"history"`
}

type History struct {
	Lat string `json:lat`
	Lng string `json:lng`
}
