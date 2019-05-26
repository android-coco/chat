package chat

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context)  {
	c.HTML(http.StatusOK, "/chat/index.shtml", nil)
}



func Concat(c *gin.Context)  {
	c.HTML(http.StatusOK, "/chat/concat.shtml", nil)
}


func Createcom(c *gin.Context)  {
	c.HTML(http.StatusOK, "/chat/createcom.shtml", nil)
}

func Group(c *gin.Context)  {
	c.HTML(http.StatusOK, "/chat/group.shtml", nil)
}

func Main(c *gin.Context)  {
	c.HTML(http.StatusOK, "/chat/main.shtml", nil)
}


func Profile(c *gin.Context)  {
	c.HTML(http.StatusOK, "/chat/profile.shtml", nil)
}

func Tabmenu(c *gin.Context)  {
	c.HTML(http.StatusOK, "/chat/tabmenu.shtml", nil)
}