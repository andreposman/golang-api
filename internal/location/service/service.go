package service

import (
	"errors"
	"fmt"
	"github.com/andreposman/golang-api/internal/location/model"
	"github.com/andreposman/golang-api/internal/location/repository"
	"github.com/gin-gonic/gin"
)

var locations = repository.GetMockData()

func GetData(c *gin.Context) (*model.Location, error ){
	orderID := c.Param("order_id")
	max := c.Query("max")
	locationsHistory := len(locations.History)
	//retreive only max len off memory

	fmt.Println("orderID:", orderID)
	fmt.Println("max:", max)
	fmt.Println("Locations: ", locations)
	fmt.Println("Len Location hist: ", locationsHistory)

	if orderID != locations.OrderID{
		return nil, errors.New("OrderID " + orderID + " not found")
	}


	return &locations, nil
}

func PutData(c *gin.Context) error{
	orderID := c.Param("order_id")
	var history model.History

	fmt.Println("orderID:", orderID)
	//validate ID

	err := c.BindJSON(&history)
	if err != nil {
		return errors.New("Error in PUT")
	}

	fmt.Println(history)
	fmt.Println("New Locations: ", locations)

	return nil
}