package main

import (
	"log"
	"os"
)

var (
	INFO        *log.Logger
	WARN        *log.Logger
	ERROR       *log.Logger
	LogFilePath string
)

func init() {
	LogFilePath = "stdlib/errorAndLog/.log"
	var logFile, _ = os.OpenFile(LogFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	INFO = log.New(logFile, "INFO: ", log.LstdFlags|log.Llongfile)
	WARN = log.New(logFile, "WARN: ", log.LstdFlags|log.Llongfile)
	ERROR = log.New(logFile, "ERROR: ", log.LstdFlags|log.Llongfile)
}
