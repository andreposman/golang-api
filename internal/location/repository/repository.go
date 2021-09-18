package repository

import "github.com/andreposman/golang-api/internal/location/model"

func GetMockData() model.Location {
	var mockLocations = model.Location{
		OrderID: "001",
		History: []model.History{
			{123, 234},
			{2322, 254},
			{13432, 3334},
			{13, 576},
			{1733, 23546},
			{7873, 876},
		},
	}
	return mockLocations
}
