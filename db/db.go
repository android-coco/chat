package db

import (
	"chat/config"
	"chat/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

var db *xorm.Engine

func InitDb(configDb config.Db) (*xorm.Engine, error) {
	args := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", configDb.User, configDb.PassWd, configDb.Host, configDb.Db)
	var err error
	db, err = xorm.NewEngine(configDb.Dialect, args)
	if err != nil {
		log.Fatalf("init db err %v \n", err)
	}
	db.ShowSQL(configDb.EnableLog)
	//用于设置最大打开的连接数，默认值为0表示不限制。
	db.SetMaxOpenConns(configDb.MaxOpenConnections)
	//设置连接池的空闲数大小
	db.SetMaxIdleConns(configDb.MaxIdleConnections)
	syncTable()
	return db, err
}

func GetDb() *xorm.Engine {
	return db
}

func syncTable(){
	err := db.Sync2(&model.User{},&model.Community{})
	if err != nil {
		log.Fatalf("init db sync table err %v \n", err)
	}
}
