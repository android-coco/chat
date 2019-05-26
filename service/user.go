package service

import (
	"chat/db"
	"chat/model"
	"chat/util"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type UserService struct {
}

//注册函数
func (user *UserService) Register(
	mobile,
	plainPwd,
	nickName,
	avatar,
	sex string) (model.User, error) {
	//检测手机号码是否存在,
	db := db.GetDb()
	tmp := model.User{}
	_, err := db.Where("mobile=? ", mobile).Get(&tmp)
	if err != nil {
		return tmp, err
	}
	//如果存在则返回提示已经注册
	if tmp.Id > 0 {
		return tmp, errors.New("该手机号已经注册")
	}
	//否则拼接插入数据
	tmp.Mobile = mobile
	tmp.Avatar = avatar
	tmp.Nickname = nickName
	if sex == "" {
		sex = model.SexNukown
	}
	tmp.Sex = sex
	tmp.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	tmp.Passwd = util.MakePasswd(plainPwd, tmp.Salt)
	tmp.Createat = time.Now()
	//token 可以是一个随机数
	tmp.Token = fmt.Sprintf("%08d", rand.Int31())
	//passwd =
	//md5 加密
	//返回新用户信息

	//插入 InserOne
	_, err = db.InsertOne(&tmp)
	//前端恶意插入特殊字符
	//数据库连接操作失败
	return tmp, err
}

//登录
func (user *UserService) Login(mobile,
	plainPwd string) (model.User, error) {

	//首先通过手机号查询用户
	tmp :=model.User{}
	db := db.GetDb()
	db.Where("mobile = ?", mobile).Get(&tmp)
	//如果没有找到
	if tmp.Id==0{
		return tmp,errors.New("该用户不存在")
	}

	//查询到了比对密码
	if !util.ValidatePasswd(plainPwd,tmp.Salt,tmp.Passwd){
		return tmp,errors.New("密码不正确")
	}

	//刷新token,安全
	str := fmt.Sprintf("%d",time.Now().Unix())
	token := util.MD5Encode(str)
	tmp.Token = token
	//返回数据
	db.ID(tmp.Id).Cols("token").Update(&tmp)
	return tmp,nil
}
