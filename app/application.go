package app

import (
	"github.com/comfysweet/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Info("application has been started")
	router.Run(":8080")
}
