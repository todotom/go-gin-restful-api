package rest

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
	getThreadsHandler GetThreadsHandler
}

func ProvideController(getThreadsHandler GetThreadsHandler) Controller {
	return Controller{getThreadsHandler}
}

func (controller Controller) Run() {
	router := gin.Default()
	router.GET("/threads", controller.getThreadsHandler.Run)
	// router.POST("/threads", postThread)
	router.Run("localhost:8080")
}
