package rest

import (
	"net/http"
	"todotom/go-gin-restful-api/forum"

	"github.com/gin-gonic/gin"
)

// var threads []forum.Thread = []forum.Thread{
// 	{Id: "a", Title: "abc"},
// }

type GetThreadsHandler struct {
	service forum.GetThreadsService
}

func ProvideGetThreadsHandler(service forum.GetThreadsService) GetThreadsHandler {
	return GetThreadsHandler{service}
}

func (handler GetThreadsHandler) Run(context *gin.Context) {
	context.JSON(http.StatusOK, handler.service.Run())
}

// func getThreads(context *gin.Context) {
// 	context.JSON(http.StatusOK, threads)
// }

// func postThread(context *gin.Context) {
// 	var newThread forum.Thread
// 	if err := context.BindJSON(&newThread); err != nil {
// 		context.JSON(http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	threads = append(threads, newThread)
// 	context.JSON(http.StatusCreated, newThread)
// }
