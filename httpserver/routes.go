package httpserver

import (
	"chat/httpserver/globle"
	"github.com/gin-gonic/gin"
)

func initRoutes(router *gin.Engine) {
	router.GET("/ping", globle.Ping)
	router.GET("/version", globle.Version)
}
