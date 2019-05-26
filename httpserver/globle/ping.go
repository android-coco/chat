package globle

import (
	"chat/model"
	"chat/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, model.ApiResp{
		ErrorNo:  util.SuccessCode,
		ErrorMsg: "pong",
	})
}
