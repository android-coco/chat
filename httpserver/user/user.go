package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context)  {
	c.HTML(http.StatusOK, "/user/login.shtml", nil)
}


func Register(c *gin.Context)  {
	c.HTML(http.StatusOK, "/user/register.shtml", nil)
}
