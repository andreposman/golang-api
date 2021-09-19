package api

import (
	"github.com/andreposman/golang-api/internal/config"
	"github.com/andreposman/golang-api/internal/location/controller"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func InitAPI() {
	router := gin.Default()
	envConfig := config.GetConfig()

	log.Info("ENV CONFIG PORTS:", envConfig)

	api := router.Group("/location")
		api.GET("/:order_id", controller.Get)
		api.PUT("/:order_id", controller.Put)
		api.DELETE(":order_id", controller.Delete)

	err := router.Run(":" + envConfig.ServerPort)
	if err != nil {
		router.Run(":" + envConfig.BackupServerPort)
	}

}
