package ctrl

import (
	"chat/args"
	"chat/httpserver/chat/ctrl"
	"chat/model"
	"chat/service"
	"chat/util"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var contactService service.ContactService

func LoadFriend(c *gin.Context) {
	var arg args.ContactArg
	//如果这个用的上,那么可以直接
	err := c.ShouldBind(&arg)
	if err != nil {
		util.RespFail(c, model.ApiResp{
			ErrorNo:  http.StatusBadRequest,
			ErrorMsg: "parameter violation!",
		}, errors.New("parameter violation "))
		return
	}

	users := contactService.SearchFriend(arg.Userid)
	util.RespOK(c, model.ApiResp{
		ErrorNo: util.SuccessCode,
		Rows:    users,
		Total:   len(users),
	})
}

func LoadCommunity(c *gin.Context) {
	var arg args.ContactArg
	//如果这个用的上,那么可以直接
	err := c.ShouldBind(&arg)
	if err != nil {
		util.RespFail(c, model.ApiResp{
			ErrorNo:  http.StatusBadRequest,
			ErrorMsg: "parameter violation!",
		}, errors.New("parameter violation "))
		return
	}
	comunitys := contactService.SearchComunity(arg.Userid)
	util.RespOK(c, model.ApiResp{
		ErrorNo: util.SuccessCode,
		Rows:    comunitys,
		Total:   len(comunitys),
	})
}
func JoinCommunity(c *gin.Context) {
	var arg args.ContactArg
	//如果这个用的上,那么可以直接
	err := c.ShouldBind(&arg)
	if err != nil {
		util.RespFail(c, model.ApiResp{
			ErrorNo:  http.StatusBadRequest,
			ErrorMsg: "parameter violation!",
		}, errors.New("parameter violation "))
		return
	}
	err = contactService.JoinCommunity(arg.Userid, arg.Dstid)
	//todo 刷新用户的群组信息
	ctrl.AddGroupId(arg.Userid, arg.Dstid)
	if err != nil {
		util.RespFail(c, model.ApiResp{
			ErrorNo:  http.StatusBadRequest,
			ErrorMsg: err.Error(),
		}, err)
		return
	} else {
		util.RespOK(c, model.ApiResp{
			ErrorNo:  util.SuccessCode,
			ErrorMsg: "",
			Data:     nil,
		})
	}
}
func CreateCommunity(c *gin.Context) {
	var arg model.Community
	//如果这个用的上,那么可以直接
	err := c.ShouldBind(&arg)
	if err != nil {
		util.RespFail(c, model.ApiResp{
			ErrorNo:  http.StatusBadRequest,
			ErrorMsg: "parameter violation!",
		}, errors.New("parameter violation "))
		return
	}
	com, err := contactService.CreateCommunity(arg)
	if err != nil {
		util.RespFail(c, model.ApiResp{
			ErrorNo:  http.StatusBadRequest,
			ErrorMsg: err.Error(),
		}, err)
		return
	} else {
		util.RespOK(c, model.ApiResp{
			ErrorNo:  util.SuccessCode,
			ErrorMsg: "",
			Data:     com,
		})
	}
}

//
func Addfriend(c *gin.Context) {
	//定义一个参数结构体
	/*request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	passwd := request.PostForm.Get("passwd")
	*/
	var arg args.ContactArg

	err := c.ShouldBind(&arg)
	fmt.Print(arg)
	if err != nil {
		util.RespFail(c, model.ApiResp{
			ErrorNo:  http.StatusBadRequest,
			ErrorMsg: "parameter violation!"+err.Error(),
		}, errors.New("parameter violation "))
		return
	}
	//调用service
	err = contactService.AddFriend(arg.Userid, arg.Dstid)
	//
	if err != nil {
		util.RespFail(c, model.ApiResp{
			ErrorNo:  http.StatusBadRequest,
			ErrorMsg: err.Error(),
		}, err)
		return
	} else {
		util.RespOK(c, model.ApiResp{
			ErrorNo:  util.SuccessCode,
			ErrorMsg: "好友添加成功",
			Data:     nil,
		})
	}
}
