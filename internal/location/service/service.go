package service

import (
	"errors"
	"github.com/andreposman/golang-api/internal/location/model"
	"github.com/andreposman/golang-api/internal/location/repository"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strconv"
)

var locations = repository.GetMockData()

func GetData(c *gin.Context) (*model.Location, error ){
	var data model.Location
	var isDigit = regexp.MustCompile(`^\d+$`).MatchString

	orderID := c.Param("order_id")
	historySize := len(locations.History)
	maxLocation := int64(historySize)

	max := c.Query("max")
	if !isDigit(max) {
		max = "0"
	}

	log.Info("ORDER ID: " + orderID)
	log.Info("MAX LOCATIONS REQUESTED BY USER: " + max)


	errID := repository.ValidateID(orderID, locations)
	if errID != nil {
		return nil, errID
	}

	maxLocation, errConv := strconv.ParseInt(max, 10, 64)
		if errConv != nil {
			log.Error("Error converting " + max + " to int")
			return nil, errors.New("Error converting " + max + " to int")
	}

	if maxLocation > int64(historySize) || maxLocation <= 0 {
		log.Info("User max param is not valid, making equal to in memory length")
		maxLocation = int64(historySize)
	}

	//probably not the most effective way of getting only N amount of history
	data.OrderID = orderID
	for i := 0; i < int(maxLocation); i++ {
		//making the most recent locations in chronological order of insertion being returned first
		data.History = append([]model.History{locations.History[i]}, data.History...)
		continue
	}

	return &data, nil
}

func PutData(c *gin.Context) error{
	orderID := c.Param("order_id")
	var history model.History

	errID := repository.ValidateID(orderID, locations)
	if errID != nil {
		return errID
	}

	err := c.BindJSON(&history)
	if err != nil {
		return errors.New("error adding new locations to history")
	}

	log.Info("ADDING NEW LOCATION DATA TO ORDER ID:", orderID)
	locations.History = append(locations.History, history)

	return nil
}



func DeleteData(c *gin.Context) error{
	orderID := c.Param("order_id")

	errID := repository.ValidateID(orderID, locations)
	if errID != nil {
		return errID
	}

	log.Info("DELETING ORDER ID:", orderID)
	locations = model.Location{}

	return nil
}