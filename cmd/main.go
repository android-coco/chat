package main

import (
	"chat/config"
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
	httpserver.Run(config.GetService().Port)
}
