package config

import (
	yaml2 "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// DB 数据库配置
type Db struct {
	EnableLog          bool   `yaml:"enable_log"`
	Dialect            string `yaml:"dialect"`
	Host               string `yaml:"host"`
	User               string `yaml:"user"`
	PassWd             string `yaml:"pass"`
	Db                 string `yaml:"db"`
	MaxOpenConnections string `yaml:"max_open_connections"`
	MaxIdleConnections string `yaml:"max_idle_connections"`
}

// Service 服务端配置
type Service struct {
	Port        string `yaml:"port"`
	ServiceUrl  string `yaml:"service_url"`
	DebugMode   bool   `yaml:"debug_mode"`
	RpcUrl      string `yaml:"rpc_url"`
	MaxPageSize int    `yaml:"max_page_size"`
}

// redis
type Redis struct {
	Host   string `yaml:"host"`
	PassWd string `yaml:"pass"`
	Db     int    `yaml:"db"`
}

type LogConfig struct {
	Path string `yaml:"path"`
}

// Config 配置
type Config struct {
	Service Service   `yaml:"service"`
	DB      Db        `yaml:"db"`
	Redis   Redis     `yaml:"redis"`
	Log     LogConfig `yaml:"log"`
}

func GetDb() Db {
	return config.DB
}

func GetService() Service {
	return config.Service
}

func GetRedis() Redis {
	return config.Redis
}

func GetLog() LogConfig {
	return config.Log
}

var config Config

func InitConfig(path string) {
	pathStr, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	yamlFile, err := ioutil.ReadFile(pathStr + path)
	if err != nil {
		log.Fatalf("yamlFile.Get err   #%v ", err)
	}
	err = yaml2.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}
