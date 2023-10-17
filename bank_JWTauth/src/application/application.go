package application

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	ComponenentsInitialization()

	router.Run(":8082")
}
