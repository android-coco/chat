package util

import "fmt"
import "github.com/cihub/seelog"

const ConfigDefaultLogConfigFile = "/../config/log.xml"

var Logger seelog.LoggerInterface

func loadAppConfig(logConfigFile string) error{
	if logConfigFile == "" {
		logConfigFile = ConfigDefaultLogConfigFile
	}
	logConfigFile, err := GetAbsPath(logConfigFile)
	logger, err := seelog.LoggerFromConfigAsFile(logConfigFile)
	if err != nil {
		fmt.Println(err)
		return err
	}
	UseLogger(logger)
	return nil
}

func InitLog(logConfigFile string) error {
	DisableLog()
	return loadAppConfig(logConfigFile)
}

// DisableLog disables all library log output
func DisableLog() {
	Logger = seelog.Disabled
}

// UseLogger uses a specified seelog.LoggerInterface to output library log. // Use this func if you are using Seelog logging system in your app.
func UseLogger(newLogger seelog.LoggerInterface) {
	Logger = newLogger
}
