package controller

import (
	"github.com/andreposman/golang-api/internal/location/service"
	"github.com/gin-gonic/gin"
	"net/http"
)


func Get(c *gin.Context) {

	data, err := service.GetData(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"order_id": data.OrderID, "history": data.History})
}

func Put(c *gin.Context) {
	err := service.PutData(c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}


func Delete(c *gin.Context) {
//TODO remove by id
}

