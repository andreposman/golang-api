package api

import (
	"github.com/andreposman/golang-api/internal/location/controller"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitAPI() {
	router := gin.Default()
	serverPort := getServerPort()

	api := router.Group("/location")
	api.GET("/:order_id", controller.Get)
	api.PUT("/:order_id", controller.Put)
	api.DELETE(":order_id", controller.Delete)

	router.Run(":" + serverPort)
}

func getServerPort() string {
	var serverPort string

	err := godotenv.Load()
	if err != nil {
		serverPort = "8080"
	}

	serverPort = os.Getenv("HISTORY_SERVER_LISTEN_ADDR")

	return serverPort
}
