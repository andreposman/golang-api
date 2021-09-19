package repository

import (
	"errors"
	"github.com/andreposman/golang-api/internal/location/model"
	log "github.com/sirupsen/logrus"
)

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


func ValidateID(orderID string, locations model.Location) error{
	if orderID != locations.OrderID {
		log.Error("Error validating orderID: ", orderID)
		return errors.New("orderID " + orderID + " not found")
	}

	return nil
}