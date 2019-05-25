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
	router.Use(middleware.EnableCors([]string{"eospark.com", "blockabc.com", "localhost:8000"}))
	//静态文件
	router.Static("/static", "/../static")
	router.StaticFS("/static", http.Dir("/../static"))
	router.StaticFile("/favicon.ico", "/../static/images/api.png")
	router.LoadHTMLGlob("/../view/**/*")
	initRoutes(router)
	err := router.Run(serverAddr)
	if err != nil {
		log.Fatalf("web server init err: %v", err)
	}
}
