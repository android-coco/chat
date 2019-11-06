package httpserver

import (
	"chat/httpserver/chat"
	chatCtrl "chat/httpserver/chat/ctrl"
	contactCtrl "chat/httpserver/contact/ctrl"
	"chat/httpserver/globle"
	"chat/httpserver/user"
	"chat/httpserver/user/ctrl"
	"github.com/gin-gonic/gin"
	"net/http"
)

func initRoutes(router *gin.Engine) {
	router.GET("/ping", globle.Ping)
	router.GET("/version", globle.Version)

	//静态文件
	router.StaticFS("/static", http.Dir("../static"))
	router.StaticFS("/mnt", http.Dir("./mnt"))
	router.StaticFile("/favicon.ico", "../static/logo.png")
	router.LoadHTMLGlob("../view/**/*")

	router.GET("/user/login.shtml", user.Login)

	router.GET("/user/register.shtml", user.Register)

	router.GET("/chat/concat.shtml", chat.Concat)

	router.GET("/chat/createcom.shtml", chat.Createcom)

	router.GET("/chat/group.shtml", chat.Group)

	router.GET("/chat/index.shtml", chat.Index)

	router.GET("/chat/main.shtml", chat.Main)

	router.GET("/chat/profile.shtml", chat.Profile)

	router.GET("/chat/tabmenu.shtml", chat.Tabmenu)

	//绑定请求和处理函数
	router.POST("/user/login", ctrl.UserLogin)
	router.POST("/user/register", ctrl.UserRegister)

	router.POST("/contact/loadcommunity", contactCtrl.LoadCommunity)
	router.POST("/contact/loadfriend", contactCtrl.LoadFriend)
	router.POST("/contact/createcommunity", contactCtrl.CreateCommunity)
	router.POST("/contact/joincommunity", contactCtrl.JoinCommunity)
	router.POST("/contact/addfriend", contactCtrl.Addfriend)

	router.GET("/chat", chatCtrl.Chat)
	router.POST("/chat", chatCtrl.Chat)
	router.POST("/attach/upload", globle.Upload)

}
