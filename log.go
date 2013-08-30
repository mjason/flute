package flute

import (
	seelog "github.com/cihub/seelog"
	"log"
)

var Logger seelog.LoggerInterface

func initAppConfig() {
	logger, err := seelog.LoggerFromConfigAsFile("./config/log.xml")
	if err != nil {
		log.Fatalln(err)
	}
	Logger = logger
}

func disableLog() {
	Logger = seelog.Disabled
}

func StartLog() {
	initAppConfig()
	disableLog()
}
