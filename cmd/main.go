package main

import (
	"chat/config"
	"chat/db"
	"chat/httpserver"
	"chat/util"
	"log"
)

func main() {
	config.InitConfig("/../config/config.yaml")
	err := util.InitLog(config.GetLog().Path)
	if nil != err {
		log.Fatalf("log init fial err:%v", err)
	}
	engine, err := db.InitDb(config.GetDb())
	if err != nil {
		log.Fatalf("init db err %v \n", err)
	}
	defer engine.Close()

	httpserver.Run(config.GetService().Port)
}
