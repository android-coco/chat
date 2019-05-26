package model

import "time"

const  (
	SexWoman  = "W"
	SexMan    = "M"
	SexNukown = "U"
)

type User struct {
	//用户ID
	Id int64 `xorm:"pk autoincr bigint(64)" form:"id" json:"id"`
	//手机号码
	Mobile string `xorm:"varchar(20)" form:"mobile" json:"mobile"`
	// 用户密码=f(plainpwd+salt)  md5
	Passwd string `xorm:"varchar(40)" form:"passwd" json:"-"`
	Avatar string `xorm:"varchar(150)" form:"avatar" json:"avatar"`
	// 性别
	Sex string `xorm:"varchar(2)" form:"sex" json:"sex"`
	// 昵称
	Nickname string `xorm:"varchar(20)" form:"nickname" json:"nickname"`
	// 随机数
	Salt string `xorm:"varchar(10)" form:"salt" json:"-"`
	//是否在线
	Online int `xorm:"int(10)" form:"online" json:"online"`
	// 用户token 鉴权用
	Token string `xorm:"varchar(40)" form:"token" json:"token"`
	Memo  string `xorm:"varchar(140)" form:"memo" json:"memo"`
	// 统计增量
	Createat time.Time `xorm:"datetime" form:"createat" json:"createat"`
}
