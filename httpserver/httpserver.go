package httpserver

import (
	"chat/config"
	"chat/httpserver/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

func Run(serverAddr string) {
	router := gin.New()
	if !config.GetService().DebugMode {
		gin.SetMode(gin.ReleaseMode)
	}
	//各种中间件
	router.Use(gin.Recovery())
	router.Use(gin.ErrorLogger())
	router.Use(middleware.EnableCors([]string{"127.0.0.1:8181","localhost:8181"}))
	initRoutes(router)
	err := router.Run(serverAddr)
	if err != nil {
		log.Fatalf("web server init err: %v", err)
	}
}
