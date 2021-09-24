package controller

import (
	"github.com/andreposman/golang-api/internal/location/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Get(c *gin.Context) {
	log.Info("GET REQUEST MADE TO: /location/{order_id}?max=<N>")

	data, err := service.GetData(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order_id": data.OrderID, "history": data.History})
}

func Put(c *gin.Context) {
	log.Info("PUT REQUEST MADE TO: location/{order_id}")

	err := service.PutData(c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}


func Delete(c *gin.Context) {
	log.Info("DELETE REQUEST MADE TO:  /location/{order_id}")

	errDelete := service.DeleteData(c)
	if errDelete != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": errDelete.Error()})
		return
	}

	c.Status(http.StatusOK)
}

