package service

import (
	"chat/db"
	"chat/model"
	"errors"
	"log"
	"time"
)

type ContactService struct {
}

//自动添加好友
func (service *ContactService) AddFriend(
	userid, //用户id 10086,
	dstid int64) error {
	//如果加自己
	if userid == dstid {
		return errors.New("不能添加自己为好友啊")
	}
	//判断是否已经加了好友
	tmp := model.Contact{}
	//查询是否已经是好友
	// 条件的链式操作
	x, err := db.GetDb().Where("ownerid = ?", userid).
		And("dstobj = ?", dstid).
		And("cate = ?", model.CONCAT_CATE_USER).
		Get(&tmp)
	if err != nil{
		log.Fatalln(x, err)
	}
	//获得1条记录
	//count()
	//如果存在记录说明已经是好友了不加
	if tmp.Id > 0 {
		return errors.New("该用户已经被添加过啦")
	}
	//事务,
	session := db.GetDb().NewSession()
	err = session.Begin()
	if err != nil {
		log.Fatalln(err)
	}
	//插自己的
	_, e2 := session.InsertOne(model.Contact{
		Ownerid:  userid,
		Dstobj:   dstid,
		Cate:     model.CONCAT_CATE_USER,
		Createat: time.Now(),
	})
	//插对方的
	_, e3 := session.InsertOne(model.Contact{
		Ownerid:  dstid,
		Dstobj:   userid,
		Cate:     model.CONCAT_CATE_USER,
		Createat: time.Now(),
	})
	//没有错误
	if e2 == nil && e3 == nil {
		//提交
		err := session.Commit()
		if err != nil {
			log.Fatalln(err)
		}
		return nil
	} else {
		//回滚
		err := session.Rollback()
		if err != nil {
			log.Fatalln(err)
		}
		if e2 != nil {
			return e2
		} else {
			return e3
		}
	}
}

func (service *ContactService) SearchComunity(userId int64) ([]model.Community) {
	conconts := make([]model.Contact, 0)
	comIds := make([]int64, 0)

	err := db.GetDb().Where("ownerid = ? and cate = ?", userId, model.CONCAT_CATE_COMUNITY).Find(&conconts)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range conconts {
		comIds = append(comIds, v.Dstobj);
	}
	coms := make([]model.Community, 0)
	if len(comIds) == 0 {
		return coms
	}
	err = db.GetDb().In("id", comIds).Find(&coms)
	if err != nil {
		log.Fatal(err)
	}
	return coms
}
func (service *ContactService) SearchComunityIds(userId int64) (comIds []int64) {
	//todo 获取用户全部群ID
	conconts := make([]model.Contact, 0)
	comIds = make([]int64, 0)

	err := db.GetDb().Where("ownerid = ? and cate = ?", userId, model.CONCAT_CATE_COMUNITY).Find(&conconts)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range conconts {
		comIds = append(comIds, v.Dstobj);
	}
	return comIds
}

//加群
func (service *ContactService) JoinCommunity(userId, comId int64) error {
	cot := model.Contact{
		Ownerid: userId,
		Dstobj:  comId,
		Cate:    model.CONCAT_CATE_COMUNITY,
	}
	_, err := db.GetDb().Get(&cot)
	if err != nil {
		log.Fatal(err)
	}
	if cot.Id == 0 {
		cot.Createat = time.Now()
		_, err := db.GetDb().InsertOne(cot)
		return err
	} else {
		return nil
	}

}

//建群
func (service *ContactService) CreateCommunity(comm model.Community) (ret model.Community, err error) {
	if len(comm.Name) == 0 {
		err = errors.New("缺少群名称")
		return ret, err
	}
	if comm.Ownerid == 0 {
		err = errors.New("请先登录")
		return ret, err
	}
	com := model.Community{
		Ownerid: comm.Ownerid,
	}
	num, err := db.GetDb().Count(&com)

	if num > 5 {
		err = errors.New("一个用户最多只能创见5个群")
		return com, err
	} else {
		comm.Createat = time.Now()
		session := db.GetDb().NewSession()
		err := session.Begin()
		if err != nil {
			log.Fatalln(err)
		}
		_, err = session.InsertOne(&comm)
		if err != nil {
			err = session.Rollback()
			if err != nil {
				log.Fatalln(err)
			}
			return com, err
		}
		_, err = session.InsertOne(
			model.Contact{
				Ownerid:  comm.Ownerid,
				Dstobj:   comm.Id,
				Cate:     model.CONCAT_CATE_COMUNITY,
				Createat: time.Now(),
			})
		if err != nil {
			err = session.Rollback()
			if err != nil {
				log.Fatalln(err)
			}
		} else {
			err := session.Commit()
			if err != nil {
				log.Fatalln(err)
			}
		}
		return com, err
	}
}

//查找好友
func (service *ContactService) SearchFriend(userId int64) ([]model.User) {
	conconts := make([]model.Contact, 0)
	objIds := make([]int64, 0)
	err := db.GetDb().Where("ownerid = ? and cate = ?", userId, model.CONCAT_CATE_USER).Find(&conconts)
	if err != nil {
		log.Fatalln(err)
	}
	for _, v := range conconts {
		objIds = append(objIds, v.Dstobj);
	}
	coms := make([]model.User, 0)
	if len(objIds) == 0 {
		return coms
	}
	err = db.GetDb().In("id", objIds).Find(&coms)
	if err != nil {
		log.Fatalln(err)
	}
	return coms
}
