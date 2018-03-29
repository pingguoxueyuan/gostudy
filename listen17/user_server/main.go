package main

import (
	"github.com/pingguoxueyuan/gostudy/listen17/log"
)

func main() {
	/*
		file := log.NewFileLog("c:/a.log")
		file.LogDebug("this is a debug log")
		file.LogWarn("this is a warn log")
	*/

	/*
		console := log.NewConsoleLog("xxxx")
		console.LogConsoleDebug("this is a console log")
		console.LogConsoleWarn("this is a warn log")
	*/
	//log := log.NewFileLog("c:/a.log")
	log := log.NewConsoleLog("xxxx")
	log.LogDebug("this is a debug file")
	log.LogWarn("this is a warn log")
}
