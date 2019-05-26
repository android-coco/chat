package ctrl

import (
	"chat/model"
	"chat/service"
	"chat/util"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var userService service.UserService

func UserLogin(c *gin.Context) {
	mobile := c.Query("mobile")
	passWd := c.Query("passwd")
	if mobile == "" || passWd == "" {
		util.RespFail(c, model.ApiResp{
			ErrorNo:  http.StatusBadRequest,
			ErrorMsg: "parameter violation!",
		}, errors.New("parameter violation "))
		return
	}

	user, err := userService.Login(mobile, passWd)
	if err != nil {
		util.RespFail(c, model.ApiResp{
			ErrorNo:  http.StatusInternalServerError,
			ErrorMsg: "login err",
		}, err)
		return
	}

	util.RespOK(c, model.ApiResp{
		ErrorNo:  util.SuccessCode,
		ErrorMsg: "",
		Data:     user,
	})

}

type userParam struct {
	Mobile   string `json:"mobile" from:"mobile" `
	Passwd   string `json:"passwd" form:"passwd"  `
	Avatar   string `json:"avatar" form:"avatar"`
	Sex      string `json:"sex" form:"sex"`
	Nickname string `json:"nickname" form:"nickname"`
}

func UserRegister(c *gin.Context) {
	var userParam userParam
	err := c.ShouldBind(&userParam)
	if nil != err{
		util.RespFail(c, model.ApiResp{
			ErrorNo:  http.StatusBadRequest,
			ErrorMsg: "parameter violation!",
		}, errors.New("parameter violation "))
		return
	}
	user, err := userService.Register(userParam.Mobile, userParam.Passwd, userParam.Nickname, userParam.Avatar, userParam.Sex)
	if err != nil {
		util.RespFail(c, model.ApiResp{
			ErrorNo:  http.StatusBadRequest,
			ErrorMsg: err.Error(),
		}, err)
		return
	}
	util.RespOK(c, model.ApiResp{
		ErrorNo:  util.SuccessCode,
		ErrorMsg: "",
		Data:     user,
	})
}
