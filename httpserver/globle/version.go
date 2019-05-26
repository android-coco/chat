package globle

import (
	"chat/model"
	"chat/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

const version = "chat_server_v1.0.0"

func Version(c *gin.Context) {
	c.JSON(http.StatusOK, model.ApiResp{
		ErrorNo:  util.SuccessCode,
		ErrorMsg: "",
		Data:     version,
	})
}
