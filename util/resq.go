package util

import (
	"chat/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespFail(c *gin.Context, resp model.ApiResp, err error) {
	Logger.Errorf("[resp fail] %s--%+v", resp.ErrorMsg, err)
	c.JSON(http.StatusOK, resp)
}

func RespOK(c *gin.Context, resp model.ApiResp) {
	Logger.Infof("[resp ok] %+v", resp.Data)
	c.JSON(http.StatusOK, resp)
}
